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
	key, _ := crypto.HexToECDSA("ad45ffdf15ad48ad0e6c4769a93338c08884f7d6bcb3879066bed489f837d1ba")
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
	task        *task
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

	roninListener := newMockListener(s.sim, context.Background(), s.config, s.utilWrapper, s.store)
	r := newMockRoninTask(s.sim, roninListener, gormDB, s.utilWrapper)
	s.task = newTask(r.listener, r.client, r.store, r.chainId, r.contracts, defaultMaxTry, VOTE_BRIDGE_OPERATORS_TASK, r.releaseTasksCh, r.util)

}

func (s *VoteBridgeOperatorsSuite) TestVoteBridgeOperatorsBySignatureSuccess() {
	defer s.sqlDB.Close()
	data := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	doneTasks, processingTasks, failedTasks, _ := s.task.voteBridgeOperatorsBySignature(&bridgeModels.Task{
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
	s.sim.Commit()

	s.Equal(len(doneTasks), 1)
	s.Equal(len(processingTasks), 1)
	s.Equal(len(failedTasks), 0)

	// Get back the signatures to validate to tx
	signatures, err := s.roninGovernance.GetBridgeOperatorVotingSignatures(&bind.CallOpts{
		Pending:     false,
		From:        s.address,
		BlockNumber: nil,
		Context:     context.Background(),
	}, common.Big1, []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	})
	s.Nil(err)
	s.Equal(len(signatures), 2)

	expected := gateway.SignatureConsumerSignature{
		V: 28,
		R: [32]byte{1, 171, 188, 209, 13, 186, 74, 245, 70, 36, 19, 117, 63, 191, 153, 154, 207, 130, 125, 248, 13, 250, 51, 93, 232, 96, 47, 115, 130, 215, 253, 206},
		S: [32]byte{26, 224, 80, 166, 50, 231, 111, 165, 175, 75, 142, 62, 226, 200, 148, 112, 229, 112, 137, 87, 142, 240, 123, 221, 197, 60, 164, 40, 65, 145, 27, 234},
	}
	s.Equal(signatures[0].R[:], expected.R[:])
	s.Equal(signatures[0].S[:], expected.S[:])
	s.Equal(signatures[0].V, expected.V)

	s.Equal(signatures[1].R[:], expected.R[:])
	s.Equal(signatures[1].S[:], expected.S[:])
	s.Equal(signatures[1].V, expected.V)
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
	task        *task
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
		GOVERNANCE_CONTRACT:           roninGovernanceAddress.Hex(),
		TRUSTED_ORGANIZATION_CONTRACT: roninTrustedOrgAddress.Hex(),
		ETH_GOVERNANCE_CONTRACT:       ethGovernanceAddress.Hex(),
	})

	roninListener := newMockListener(s.sim, context.Background(), s.config, s.utilWrapper, s.store)
	r := newMockRoninTask(s.sim, roninListener, gormDB, s.utilWrapper)
	s.task = newTask(r.listener, r.client, r.store, r.chainId, r.contracts, defaultMaxTry, RELAY_BRIDGE_OPERATORS_TASK, r.releaseTasksCh, r.util)

}

func (s *RelayBridgeOperatorsSuite) TestRelayBridgeOperatorsSuccess() {
	defer s.sqlDB.Close()

	// Setup operators signatures
	data := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	doneTasks, processingTasks, failedTasks, _ := s.task.voteBridgeOperatorsBySignature(&bridgeModels.Task{
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

	// Setup trusted organizations
	_, err := s.task.util.SendContractTransaction(s.task.listener.GetValidatorSign(), s.task.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return s.roninTrustedOrg.SetTrustedOrganizations(opts, []common.Address{
			common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
			common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
		})
	})
	s.Nil(err)
	s.sim.Commit()

	// Run the actual test
	data = "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	doneTasks, processingTasks, failedTasks, _ = s.task.relayBridgeOperators(&bridgeModels.Task{
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
	s.sim.Commit()

	s.Equal(len(doneTasks), 1)
	s.Equal(len(processingTasks), 1)
	s.Equal(len(failedTasks), 0)

	// Get the result from smart contract to verify again
	signatures, err := s.ethGovernance.GetBridgeOperatorVotingSignatures(nil, common.Big1, []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	})
	expected := gateway.SignatureConsumerSignature{
		V: 28,
		R: [32]byte{1, 171, 188, 209, 13, 186, 74, 245, 70, 36, 19, 117, 63, 191, 153, 154, 207, 130, 125, 248, 13, 250, 51, 93, 232, 96, 47, 115, 130, 215, 253, 206},
		S: [32]byte{26, 224, 80, 166, 50, 231, 111, 165, 175, 75, 142, 62, 226, 200, 148, 112, 229, 112, 137, 87, 142, 240, 123, 221, 197, 60, 164, 40, 65, 145, 27, 234},
	}

	s.Nil(err)
	s.Equal(len(signatures), 2)

	s.Equal(signatures[0].R[:], expected.R[:])
	s.Equal(signatures[0].S[:], expected.S[:])
	s.Equal(signatures[0].V, expected.V)

	s.Equal(signatures[1].R[:], expected.R[:])
	s.Equal(signatures[1].S[:], expected.S[:])
	s.Equal(signatures[1].V, expected.V)
}

type CommonTestSuite struct {
	suite.Suite
	SimulatedSuite
	roninTask       *MockRoninTask
	task            *task
	roninGovernance *contract_tests.RoninGovernance
}

func TestRunCommonSuite(t *testing.T) {
	suite.Run(t, new(CommonTestSuite))
}

func (s *CommonTestSuite) SetupTest() {
	s.SimulatedSuite.SetupTest()

	gormDB, sqlDB, _, _ := newMockDB()
	defer sqlDB.Close()

	address, _, contract, err := contract_tests.DeployRoninGovernance(s.auth, s.sim)
	s.roninGovernance = contract
	s.Nil(err)
	s.sim.Commit()

	utilWrapper := utils.NewUtils()
	store := stores.NewMainStore(gormDB)
	config := newMockConfig(s.privateKey, map[string]string{
		GOVERNANCE_CONTRACT: address.Hex(),
	})

	roninListener := newMockListener(s.sim, context.Background(), config, utilWrapper, store)
	s.roninTask = newMockRoninTask(s.sim, roninListener, gormDB, utilWrapper)
	s.task = newTask(s.roninTask.listener, s.roninTask.client, s.roninTask.store, s.roninTask.chainId, s.roninTask.contracts, defaultMaxTry, RELAY_BRIDGE_OPERATORS_TASK, s.roninTask.releaseTasksCh, s.roninTask.util)
}

func (s *CommonTestSuite) TestUnpackBridgeOperatorsApprovedEventSuccess() {
	data := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	event, err := s.task.unpackBridgeOperatorsApprovedEvent(&bridgeModels.Task{
		Data: data,
	})

	expectedPeriod := common.Big1
	expectedBridgeOperators := []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	}
	s.Equal(expectedPeriod, event.Period)
	s.Equal(expectedBridgeOperators, event.Operators)
	s.Nil(err)
}

func (s *CommonTestSuite) TestUnpackBridgeOperatorsUpdatedEventSuccess() {
	data := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b6bc5bc0410773a3f86b1537ce7495c52e38f88b0000000000000000000000004a4bc674a97737376cfe990ae2fe0d2b6e738393"
	event, err := s.task.unpackBridgeOperatorSetUpdatedEvent(&bridgeModels.Task{
		Data: data,
	})

	expectedPeriod := common.Big1
	expectedBridgeOperators := []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	}
	s.Equal(expectedPeriod, event.Period)
	s.Equal(expectedBridgeOperators, event.BridgeOperators)
	s.Nil(err)
}

func (s *CommonTestSuite) TestSignBridgeOperatorsBallotSuccess() {
	period := int64(9263652)
	bridgeOperators := []interface{}{
		"0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B",
		"0x4a4bc674A97737376cFE990aE2fE0d2B6E738393",
	}

	hash, err := signBridgeOperatorsBallot(&signDataOpts{
		SignTypedDataCallback: func(typedData core2.TypedData) (hexutil.Bytes, error) {
			return s.roninTask.util.SignTypedData(typedData, s.roninTask.listener.GetValidatorSign())
		},
	}, period, bridgeOperators)
	sig := parseSignatureAsRsv(hash)
	fmt.Println("hash", common.BytesToHash(hash))
	fmt.Println("r", common.Bytes2Hex(sig.R[:]))
	fmt.Println("s", common.Bytes2Hex(sig.S[:]))
	fmt.Println("v", sig.V)
	expected := gateway.SignatureConsumerSignature{
		V: 27,
		R: [32]byte{207, 155, 162, 143, 111, 90, 47, 92,
			106, 236, 250, 180, 130, 207, 27, 167,
			223, 254, 139, 65, 232, 17, 158, 68,
			173, 175, 22, 57, 53, 248, 65, 194},
		S: [32]byte{84, 81, 183, 143, 10, 189, 8, 112,
			89, 210, 207, 134, 196, 25, 141, 170,
			110, 91, 108, 244, 140, 163, 80, 245,
			229, 200, 172, 168, 57, 94, 116, 150},
	}
	s.Equal(sig.R[:], expected.R[:])
	s.Equal(sig.S[:], expected.S[:])
	s.Equal(sig.V, expected.V)
	s.Nil(err)
}

func (s *CommonTestSuite) TestCallVoteBridgeOperatorsBySignaturesSuccess() {
	//client, _ := ethclient.Dial("http://localhost:8545")
	//privateKey, _ := crypto.HexToECDSA("ad45ffdf15ad48ad0e6c4769a93338c08884f7d6bcb3879066bed489f837d1ba")

	typedData := core2.TypedData{
		Types: core2.Types{
			"EIP712Domain": []core2.Type{
				{
					Name: "name", Type: "string",
				},
				{
					Name: "version", Type: "string",
				},
				{
					Name: "salt", Type: "string",
				},
			},
			"BridgeOperatorsBallot": []core2.Type{
				{
					Name: "period", Type: "uint256",
				},
				{
					Name: "operators", Type: "address[]",
				},
			},
		},
		PrimaryType: "BridgeOperatorsBallot",
		Domain: core2.TypedDataDomain{
			Name:    "GovernanceAdmin",
			Version: "1",
			Salt:    "0xe3922a0bff7e80c6f7465bc1b150f6c95d9b9203f1731a09f86e759ea1eaa306",
		},
		Message: core2.TypedDataMessage{
			"period": math.NewHexOrDecimal256(9263652),
			"operators": []interface{}{
				"0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B",
				"0x4a4bc674A97737376cFE990aE2fE0d2B6E738393",
			},
		},
	}
	b, err := s.roninTask.util.SignTypedData(typedData, s.roninTask.listener.GetValidatorSign())
	fmt.Println("err", err)
	fmt.Println("sig", common.Bytes2Hex(b))

	// 0x520ba520cf2ef7d0adfa6dfdf93c983f88dc0d8a05ed58a620fb1acc098e0ebd136510aaf35132cf34e773ac9fe6e32c0dd6b6ca5ee052144bf7f4391cdf2fcb1b

	//domainSeparator, _ := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	//fmt.Println("domainSeparator", domainSeparator)
	//typedDataHash, _ := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	//rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	//fmt.Println("rawData", rawData)
	//
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(2021))
	//nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	//gasPrice, _ := client.SuggestGasPrice(context.Background())
	//transactor, _ := roninGovernance.NewGovernanceTransactor(common.HexToAddress("0xFBD674ba8E716F431041e2b4195CF4f1975DdFB5"), client)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)     // in wei
	//auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = gasPrice
	//
	//tx, err := transactor.VoteBridgeOperatorsBySignatures(auth, big.NewInt(9263260), []common.Address{
	//	common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
	//	common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	//}, []roninGovernance.SignatureConsumerSignature{
	//	{
	//		V: 27,
	//		R: [32]byte{234, 4, 6, 171, 43, 100, 120, 8, 161, 178, 15, 9, 116, 119, 92, 176, 75, 244, 43, 197, 215, 90, 5, 115, 211,
	//			23, 97, 158, 227, 85, 116, 120},
	//		S: [32]byte{46, 107, 139, 81, 45, 73, 206, 154, 20, 252, 92, 115, 171, 56, 64, 218, 115, 100, 201, 125, 39, 238, 40, 63,
	//			77, 226, 24, 177, 235, 209, 21, 31},
	//	},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(tx.Hash().Hex())
	//receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	//fmt.Println(receipt.Status)
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
