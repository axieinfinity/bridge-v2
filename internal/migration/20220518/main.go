package models

import (
	"fmt"
	"github.com/axieinfinity/bridge-v2/generated_contracts/ronin/gateway"
	models "github.com/axieinfinity/bridge-v2/internal/migration/20220515"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"math/big"
	"reflect"
)

func getCaller(cfg *types.Config) (*gateway.GatewayCallerRaw, error) {
	u := &utils.Utils{}
	client, _ := u.NewEthClient(cfg.Listeners["Ronin"].RpcUrl)
	caller, err := gateway.NewGatewayCaller(common.HexToAddress(cfg.Listeners["Ronin"].Contracts[types.GATEWAY_CONTRACT]), client)
	if err != nil {
		return nil, err
	}
	return &gateway.GatewayCallerRaw{Contract: caller}, nil
}

func getWithdrawalQuantity(caller *gateway.GatewayCallerRaw, withdrawalId *big.Int) (*big.Int, error) {
	var results []interface{}
	if err := caller.Call(nil, &results, "withdrawal", withdrawalId); err != nil {
		return nil, err
	}
	if len(results) != 5 {
		return nil, fmt.Errorf("mismatched fields expected 5 got %d", len(results))
	}
	info := reflect.ValueOf(results[4])
	var (
		quantity *big.Int
		ok       bool
	)
	if quantity, ok = info.Field(2).Interface().(*big.Int); ok {
		return quantity, nil
	}
	return nil, fmt.Errorf("cannot cast quantity at type:%s to *big.Int", reflect.TypeOf(quantity).String())
}

func Migrate(cfg *types.Config) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220518",
		Migrate: func(tx *gorm.DB) error {
			caller, err := getCaller(cfg)
			if err != nil {
				return err
			}
			var withdrawals []*Withdrawal
			// loop through database get withdrawalId and query their quantity and update
			rows, err := tx.Model(&models.Withdrawal{}).Order("id ASC").Rows()
			if err != nil {
				return err
			}
			defer rows.Close()

			for rows.Next() {
				var withdrawal models.Withdrawal
				if err = tx.ScanRows(rows, &withdrawal); err != nil {
					return err
				}
				// get quantity from smart contract based on withdrawal_id
				quantity, err := getWithdrawalQuantity(caller, big.NewInt(withdrawal.WithdrawalId))
				if err != nil {
					return err
				}
				withdrawals = append(withdrawals, &Withdrawal{
					ID:                   withdrawal.ID,
					WithdrawalId:         withdrawal.WithdrawalId,
					ExternalAddress:      withdrawal.ExternalAddress,
					ExternalTokenAddress: withdrawal.ExternalTokenAddress,
					ExternalChainId:      withdrawal.ExternalChainId,
					RoninAddress:         withdrawal.RoninAddress,
					RoninTokenAddress:    withdrawal.RoninTokenAddress,
					TokenErc:             withdrawal.TokenErc,
					TokenId:              withdrawal.TokenId,
					TokenQuantity:        quantity.String(),
					Transaction:          "",
				})
			}
			// drop current table
			if err = tx.Migrator().DropTable("withdrawal"); err != nil {
				return err
			}
			if err = tx.AutoMigrate(&Withdrawal{}); err != nil {
				return err
			}
			// loop through withdrawals and insert all data to new table
			return tx.Model(&Withdrawal{}).CreateInBatches(withdrawals, 100).Error
		},
		Rollback: func(tx *gorm.DB) error {
			var withdrawals []*models.Withdrawal
			// loop through database get withdrawalId and query their quantity and update
			rows, err := tx.Model(&Withdrawal{}).Order("id ASC").Rows()
			if err != nil {
				return err
			}
			defer rows.Close()
			for rows.Next() {
				var withdrawal Withdrawal
				if err = tx.ScanRows(rows, &withdrawal); err != nil {
					return err
				}
				quantity, ok := big.NewInt(0).SetString(withdrawal.TokenQuantity, 10)
				if !ok {
					return fmt.Errorf("cannot cast quantity to *big.Int")
				}
				withdrawals = append(withdrawals, &models.Withdrawal{
					ID:                   withdrawal.ID,
					WithdrawalId:         withdrawal.WithdrawalId,
					ExternalAddress:      withdrawal.ExternalAddress,
					ExternalTokenAddress: withdrawal.ExternalTokenAddress,
					ExternalChainId:      withdrawal.ExternalChainId,
					RoninAddress:         withdrawal.RoninAddress,
					RoninTokenAddress:    withdrawal.RoninTokenAddress,
					TokenErc:             withdrawal.TokenErc,
					TokenId:              withdrawal.TokenId,
					TokenQuantity:        quantity.Int64(),
				})
			}
			// drop new table and replace with old one
			// drop current table
			if err = tx.Migrator().DropTable("withdrawal"); err != nil {
				return err
			}
			if err = tx.AutoMigrate(&models.Withdrawal{}); err != nil {
				return err
			}
			// loop through withdrawals and insert all data to new table
			return tx.Model(&models.Withdrawal{}).CreateInBatches(withdrawals, 100).Error
		},
	}
}
