package task

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	roninGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/governance"
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
	"github.com/ethereum/go-ethereum/ethclient"
	core2 "github.com/ethereum/go-ethereum/signer/core"
	"github.com/stretchr/testify/suite"
	"log"
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
	key, _ := crypto.HexToECDSA("927a004a8a0e854813d4950517551eb5b0eb87a82e466438dcbc1b906572b125")
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
	period := int64(1)
	bridgeOperators := []interface{}{
		common.HexToAddress("0xf6fd5FcA4Bd769BA495B29B98dba5F2eCF4CEED3").Hex(),
	}

	hash, err := signBridgeOperatorsBallot(&signDataOpts{
		ChainId:           math.NewHexOrDecimal256(s.roninTask.chainId.Int64()),
		VerifyingContract: s.roninTask.contracts[GOVERNANCE_CONTRACT],
		SignTypedDataCallback: func(typedData core2.TypedData) (hexutil.Bytes, error) {
			return s.roninTask.util.SignTypedData(typedData, s.roninTask.listener.GetValidatorSign())
		},
	}, period, bridgeOperators)
	sig := parseSignatureAsRsv(hash)
	expected := gateway.SignatureConsumerSignature{
		V: 28,
		R: [32]byte{34, 85, 150, 68, 71, 72, 216, 50, 144, 16, 124, 240, 132, 182, 39, 100, 171, 165, 116, 88, 19, 157, 177, 176, 236, 136, 93, 46, 98, 79, 58, 121},
		S: [32]byte{124, 127, 93, 159, 169, 72, 75, 170, 233, 137, 56, 201, 209, 119, 2, 223, 117, 127, 93, 2, 69, 101, 21, 43, 102, 52, 44, 28, 146, 170, 177, 138},
	}
	s.Equal(sig.R[:], expected.R[:])
	s.Equal(sig.S[:], expected.S[:])
	s.Equal(sig.V, expected.V)
	s.Nil(err)
}

func (s *CommonTestSuite) TestCallVoteBridgeOperatorsBySignaturesSuccess() {
	client, _ := ethclient.Dial("http://localhost:8545")
	privateKey, _ := crypto.HexToECDSA("")

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(2021))

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transactor, _ := roninGovernance.NewGovernanceTransactor(common.HexToAddress("0xFBD674ba8E716F431041e2b4195CF4f1975DdFB5"), client)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := transactor.VoteBridgeOperatorsBySignatures(auth, big.NewInt(9263210), []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	}, []roninGovernance.SignatureConsumerSignature{
		{
			V: 27,
			R: [32]byte{59, 189, 60, 136, 128, 85, 64, 82, 35, 23, 82, 186, 102, 135, 98, 112, 20, 240, 39, 251, 94, 21, 150, 18, 211, 6, 82, 6, 135, 5, 166, 163},
			S: [32]byte{90, 221, 107, 82, 153, 157, 224, 54, 243, 122, 163, 68, 13, 81, 174, 96, 80, 20, 168, 21, 17, 67, 28, 67, 209, 174, 82, 159, 87, 115, 82, 112},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
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
