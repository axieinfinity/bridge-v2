package task

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	internal "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/stores"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/contract_tests"
	bridgeModels "github.com/axieinfinity/bridge-v2/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	core2 "github.com/ethereum/go-ethereum/signer/core"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

type SimulatedSuite struct {
	privateKey string
	auth       *bind.TransactOpts
	address    common.Address
	gAlloc     core.GenesisAlloc
	sim        *backends.SimulatedBackend
}

func (s *SimulatedSuite) SetupTest() {
	key, _ := crypto.GenerateKey()
	s.privateKey = fmt.Sprintf("%x", crypto.FromECDSA(key))
	s.auth = bind.NewKeyedTransactor(key)
	s.address = s.auth.From
	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.address: {
			Balance: big.NewInt(10_000_000_000),
		},
	}
	s.sim = backends.NewSimulatedBackend(s.gAlloc, 1_000_000)
}

type VoteBridgeOperatorsSuite struct {
	suite.Suite
	SimulatedSuite
	roninGovernance *contract_tests.RoninGovernance

	utilWrapper utils.Utils
	store       stores.MainStore
	config      *internal.LsConfig
	sqlDB       *sql.DB
	mockDB      sqlmock.Sqlmock
}

func TestRunVoteBridgeOperatorsSuite(t *testing.T) {
	suite.Run(t, new(VoteBridgeOperatorsSuite))
}

func (s *VoteBridgeOperatorsSuite) SetupTest() {
	s.SimulatedSuite.SetupTest()

	address, _, contract, err := contract_tests.DeployRoninGovernance(s.auth, s.sim)
	s.roninGovernance = contract
	s.Nil(err)
	s.sim.Commit()

	gormDB, sqlDB, mockDB, _ := newMockDB()
	s.sqlDB = sqlDB
	s.mockDB = mockDB

	s.utilWrapper = utils.NewUtils()
	s.store = stores.NewMainStore(gormDB)
	s.config = newMockConfig(s.privateKey, map[string]string{
		GATEWAY_CONTRACT:    common.BytesToAddress([]byte{100}).Hex(),
		GOVERNANCE_CONTRACT: address.Hex(),
	})
}

func (s *VoteBridgeOperatorsSuite) TestVoteBridgeOperatorsBySignatureSuccess() {
	gormDB, sqlDB, _, _ := newMockDB()
	defer sqlDB.Close()

	roninListener := newMockListener(s.sim, context.Background(), s.config, s.utilWrapper, s.store)
	r := newMockRoninTask(s.sim, roninListener, gormDB, s.utilWrapper)
	voteBridgeOperatorsTask := newTask(r.listener, r.client, r.store, r.chainId, r.contracts, defaultMaxTry, VOTE_BRIDGE_OPERATORS_TASK, r.releaseTasksCh, r.util)

	data := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	doneTasks, processingTasks, failedTasks, _ := voteBridgeOperatorsTask.voteBridgeOperatorsBySignature(&bridgeModels.Task{
		ChainId:         RoninChainId,
		FromChainId:     EthereumChainId,
		Type:            VOTE_BRIDGE_OPERATORS_TASK,
		Data:            data,
		Retries:         0,
		Status:          STATUS_PENDING,
		LastError:       "",
		TransactionHash: common.Bytes2Hex([]byte{100}),
		FromTransaction: common.Bytes2Hex([]byte{101}),
	})

	s.Equal(len(doneTasks), 1)
	s.Equal(len(processingTasks), 1)
	s.Equal(len(failedTasks), 0)
}

type RelayBridgeOperatorsSuite struct {
	suite.Suite
	SimulatedSuite
	roninGovernance *contract_tests.RoninGovernance
	roninTrustedOrg *contract_tests.RoninTrustedOrganization
	ethGovernance   *contract_tests.EthGovernance

	utilWrapper utils.Utils
	store       stores.MainStore
	config      *internal.LsConfig
	sqlDB       *sql.DB
	mockDB      sqlmock.Sqlmock
}

func TestRunRelayBridgeOperatorsSuite(t *testing.T) {
	suite.Run(t, new(RelayBridgeOperatorsSuite))
}

func (s *RelayBridgeOperatorsSuite) SetupTest() {
	s.SimulatedSuite.SetupTest()

	roninGovernanceAddress, _, roninGovernance, err := contract_tests.DeployRoninGovernance(s.auth, s.sim)
	s.roninGovernance = roninGovernance
	s.Nil(err)

	roninTrustedOrgAddress, _, roninTrustedOrg, err := contract_tests.DeployRoninTrustedOrganization(s.auth, s.sim)
	s.roninTrustedOrg = roninTrustedOrg
	s.Nil(err)

	ethGovernanceAddress, _, ethGovernance, err := contract_tests.DeployEthGovernance(s.auth, s.sim)
	s.ethGovernance = ethGovernance
	s.Nil(err)

	s.sim.Commit()

	gormDB, sqlDB, mockDB, _ := newMockDB()
	s.sqlDB = sqlDB
	s.mockDB = mockDB

	s.utilWrapper = utils.NewUtils()
	s.store = stores.NewMainStore(gormDB)
	s.config = newMockConfig(s.privateKey, map[string]string{
		GOVERNANCE_CONTRACT:     roninGovernanceAddress.Hex(),
		TRUSTED_ORGS_CONTRACT:   roninTrustedOrgAddress.Hex(),
		ETH_GOVERNANCE_CONTRACT: ethGovernanceAddress.Hex(),
	})
}

func (s *RelayBridgeOperatorsSuite) TestRelayBridgeOperatorsSuccess() {
	gormDB, sqlDB, _, _ := newMockDB()
	defer sqlDB.Close()

	roninListener := newMockListener(s.sim, context.Background(), s.config, s.utilWrapper, s.store)
	r := newMockRoninTask(s.sim, roninListener, gormDB, s.utilWrapper)
	relayBridgeOperatorsTask := newTask(r.listener, r.client, r.store, r.chainId, r.contracts, defaultMaxTry, RELAY_BRIDGE_OPERATORS_TASK, r.releaseTasksCh, r.util)

	// Setup trusted organizations
	_, err := r.util.SendContractTransaction(r.listener.GetValidatorSign(), r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return s.roninTrustedOrg.SetTrustedOrganizations(opts, []common.Address{
			common.BytesToAddress([]byte{200}),
			common.BytesToAddress([]byte{201}),
			common.BytesToAddress([]byte{202}),
		})
	})
	s.Nil(err)

	data := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	doneTasks, processingTasks, failedTasks, _ := relayBridgeOperatorsTask.relayBridgeOperators(&bridgeModels.Task{
		ChainId:         RoninChainId,
		FromChainId:     EthereumChainId,
		Type:            VOTE_BRIDGE_OPERATORS_TASK,
		Data:            data,
		Retries:         0,
		Status:          STATUS_PENDING,
		LastError:       "",
		TransactionHash: common.Bytes2Hex([]byte{100}),
		FromTransaction: common.Bytes2Hex([]byte{101}),
	})

	s.Equal(len(doneTasks), 1)
	s.Equal(len(processingTasks), 1)
	s.Equal(len(failedTasks), 0)
}

type CommonTestSuite struct {
	suite.Suite
	SimulatedSuite
	roninTask *MockRoninTask
}

func TestRunCommonSuite(t *testing.T) {
	suite.Run(t, new(CommonTestSuite))
}

func (s *CommonTestSuite) SetupTest() {
	s.SimulatedSuite.SetupTest()

	gormDB, sqlDB, _, _ := newMockDB()
	defer sqlDB.Close()

	utilWrapper := utils.NewUtils()
	store := stores.NewMainStore(gormDB)
	config := newMockConfig(s.privateKey, map[string]string{})

	roninListener := newMockListener(s.sim, context.Background(), config, utilWrapper, store)
	s.roninTask = newMockRoninTask(s.sim, roninListener, gormDB, utilWrapper)
}

func (s *CommonTestSuite) TestUnpackBridgeOperatorsApprovedEventSuccess() {
	data := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	event, err := unpackBridgeOperatorsApprovedEvent(&bridgeModels.Task{
		Data: data,
	})

	expectedPeriod := big.NewInt(1)
	expectedBridgeOperators := []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	}
	s.Equal(expectedPeriod, event.Period)
	s.Equal(expectedBridgeOperators, event.BridgeOperators)
	s.Nil(err)
}

func (s *CommonTestSuite) TestUnpackBridgeOperatorsUpdatedEventSuccess() {
	data := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	event, err := unpackBridgeOperatorsUpdatedEvent(&bridgeModels.Task{
		Data: data,
	})

	expectedPeriod := big.NewInt(1)
	expectedBridgeOperators := []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	}
	s.Equal(expectedPeriod, event.Period)
	s.Equal(expectedBridgeOperators, event.BridgeOperators)
	s.Nil(err)
}

func (s *CommonTestSuite) TestSignBridgeOperatorsBallotSuccess() {
	period := int64(1)
	bridgeOperators := []interface{}{
		common.HexToAddress("0xf6fd5FcA4Bd769BA495B29B98dba5F2eCF4CEED3").Hex(),
	}

	hash, err := signBridgeOperatorsBallot(&signDataOpts{
		ChainId:           math.NewHexOrDecimal256(s.roninTask.chainId.Int64()),
		VerifyingContract: s.roninTask.contracts[GATEWAY_CONTRACT],
		SignTypedDataCallback: func(typedData core2.TypedData) (hexutil.Bytes, error) {
			return s.roninTask.util.SignTypedData(typedData, s.roninTask.listener.GetValidatorSign())
		},
	}, period, bridgeOperators)
	sig := parseSignatureAsRsv(hash)
	expected := gateway.SignatureConsumerSignature{
		V: 28,
		R: [32]byte{68, 51, 43, 101, 96, 121, 145, 131,
			26, 25, 235, 92, 56, 252, 36, 99,
			87, 60, 74, 70, 70, 139, 67, 174,
			78, 162, 65, 55, 210, 49, 49, 78,
		},
		S: [32]byte{
			119, 84, 228, 229, 174, 119, 108, 82,
			71, 10, 41, 145, 158, 121, 52, 97,
			164, 243, 191, 169, 174, 187, 101, 99,
			190, 151, 76, 97, 143, 143, 18, 68,
		},
	}
	s.Equal(sig.R[:], expected.R[:])
	s.Equal(sig.S[:], expected.S[:])
	s.Equal(sig.V, expected.V)
	s.Nil(err)
}

func (s *CommonTestSuite) TestParseSignatureAsRsvSuccess() {
	hash := common.Hex2Bytes("5dabb5919c805eb67fd8e651a8bb99b1a1b7359ad64b8719bc5a83281bf171ee1349a75c125b8534176df1e290d8510370df23459f18f609cb36963596fca0ad1b")
	sig := parseSignatureAsRsv(hash)
	expected := gateway.SignatureConsumerSignature{
		V: 27,
		R: [32]byte{
			93, 171, 181, 145, 156, 128, 94,
			182, 127, 216, 230, 81, 168, 187,
			153, 177, 161, 183, 53, 154, 214,
			75, 135, 25, 188, 90, 131, 40,
			27, 241, 113, 238},
		S: [32]byte{
			19, 73, 167, 92, 18, 91, 133, 52,
			23, 109, 241, 226, 144, 216, 81, 3,
			112, 223, 35, 69, 159, 24, 246, 9,
			203, 54, 150, 53, 150, 252, 160, 173,
		},
	}
	s.Equal(sig.R[:], expected.R[:])
	s.Equal(sig.S[:], expected.S[:])
	s.Equal(sig.V, expected.V)
}
