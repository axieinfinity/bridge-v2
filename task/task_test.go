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
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"): {
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
	data := "7b2261646472657373223a22307862363863626264643463376233653465333963303335626433353666633562646237653261353065222c22746f70696373223a5b22307838643764353139653831633262386463363762343466643634356664326338383035313130643961623164363433653364643638623632326264653333316666222c22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030386436656465225d2c2264617461223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303230303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030323030303030303030303030303030303030303030303030306236626335626330343130373733613366383662313533376365373439356335326533386638386230303030303030303030303030303030303030303030303034613462633637346139373733373337366366653939306165326665306432623665373338333933222c22626c6f636b4e756d626572223a223078366164222c227472616e73616374696f6e48617368223a22307839356137303665643066623631643164333939623966376235616363396132616538396361616239626639353762326338613236303862343665623332666364222c227472616e73616374696f6e496e646578223a22307830222c22626c6f636b48617368223a22307838316462623832326435353735613334623365663562366266656334353762386235623762656335356338623461356463636365613132643861646238636137222c226c6f67496e646578223a22307836222c2272656d6f766564223a66616c73657d"
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
	}, big.NewInt(9268958), []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	})
	s.Nil(err)
	s.Equal(len(signatures), 2)

	expected := gateway.SignatureConsumerSignature{
		V: 28,
		R: [32]byte{227, 160, 125, 116, 187, 43, 58, 12, 8, 93, 241, 48, 161, 128, 31, 52, 123, 46, 251, 218, 193, 41, 92, 121, 239, 188, 198, 180, 235, 54, 125, 74},
		S: [32]byte{80, 209, 80, 160, 3, 36, 249, 77, 119, 94, 18, 119, 127, 101, 154, 10, 170, 57, 169, 120, 106, 86, 134, 120, 176, 23, 137, 108, 95, 156, 92, 136},
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

	// Setup trusted organizations
	_, err := s.task.util.SendContractTransaction(s.task.listener.GetValidatorSign(), s.task.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return s.roninTrustedOrg.SetTrustedOrganizations(opts, []common.Address{
			common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
			common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
		})
	})
	s.Nil(err)
	s.sim.Commit()

	// Setup operators signatures
	data := "7b2261646472657373223a22307862363863626264643463376233653465333963303335626433353666633562646237653261353065222c22746f70696373223a5b22307838643764353139653831633262386463363762343466643634356664326338383035313130643961623164363433653364643638623632326264653333316666222c22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030386436656465225d2c2264617461223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303230303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030323030303030303030303030303030303030303030303030306236626335626330343130373733613366383662313533376365373439356335326533386638386230303030303030303030303030303030303030303030303034613462633637346139373733373337366366653939306165326665306432623665373338333933222c22626c6f636b4e756d626572223a223078366164222c227472616e73616374696f6e48617368223a22307839356137303665643066623631643164333939623966376235616363396132616538396361616239626639353762326338613236303862343665623332666364222c227472616e73616374696f6e496e646578223a22307830222c22626c6f636b48617368223a22307838316462623832326435353735613334623365663562366266656334353762386235623762656335356338623461356463636365613132643861646238636137222c226c6f67496e646578223a22307836222c2272656d6f766564223a66616c73657d"
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

	// Run the actual test
	data = "7b2261646472657373223a22307865653731623633313231623033643739626437633035623337373432353464646437376130333661222c22746f70696373223a5b22307831353939623034613131303464313965663533346463313737663364653031363465663565346239396661643734383565646131333436303066636135663032225d2c2264617461223a2230783030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303038643665646430303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303430303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030323030303030303030303030303030303030303030303030306236626335626330343130373733613366383662313533376365373439356335326533386638386230303030303030303030303030303030303030303030303034613462633637346139373733373337366366653939306165326665306432623665373338333933222c22626c6f636b4e756d626572223a223078366131222c227472616e73616374696f6e48617368223a22307831623465366134613565613739663566373732656133366663363634663461393035323636636266343338356534663262393332346363663565373161396263222c227472616e73616374696f6e496e646578223a22307830222c22626c6f636b48617368223a22307861643937303130353634633265636332346262323366306361336432303432373033376534383632383261626564653566343931643236623837616664663038222c226c6f67496e646578223a22307830222c2272656d6f766564223a66616c73657d"
	doneTasks, processingTasks, failedTasks, _ = s.task.relayBridgeOperators(&bridgeModels.Task{
		ChainId:         RoninChainId,
		FromChainId:     EthereumChainId,
		Type:            RELAY_BRIDGE_OPERATORS_TASK,
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
	signatures, err := s.ethGovernance.GetBridgeOperatorVotingSignatures(nil, big.NewInt(9268957), []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	})
	expected := gateway.SignatureConsumerSignature{
		V: 28,
		R: [32]byte{227, 160, 125, 116, 187, 43, 58, 12, 8, 93, 241, 48, 161, 128, 31, 52, 123, 46, 251, 218, 193, 41, 92, 121, 239, 188, 198, 180, 235, 54, 125, 74},
		S: [32]byte{80, 209, 80, 160, 3, 36, 249, 77, 119, 94, 18, 119, 127, 101, 154, 10, 170, 57, 169, 120, 106, 86, 134, 120, 176, 23, 137, 108, 95, 156, 92, 136},
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
	data := "7b2261646472657373223a22307865653731623633313231623033643739626437633035623337373432353464646437376130333661222c22746f70696373223a5b22307831353939623034613131303464313965663533346463313737663364653031363465663565346239396661643734383565646131333436303066636135663032225d2c2264617461223a2230783030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303038643665646430303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303430303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030323030303030303030303030303030303030303030303030306236626335626330343130373733613366383662313533376365373439356335326533386638386230303030303030303030303030303030303030303030303034613462633637346139373733373337366366653939306165326665306432623665373338333933222c22626c6f636b4e756d626572223a223078366131222c227472616e73616374696f6e48617368223a22307831623465366134613565613739663566373732656133366663363634663461393035323636636266343338356534663262393332346363663565373161396263222c227472616e73616374696f6e496e646578223a22307830222c22626c6f636b48617368223a22307861643937303130353634633265636332346262323366306361336432303432373033376534383632383261626564653566343931643236623837616664663038222c226c6f67496e646578223a22307830222c2272656d6f766564223a66616c73657d"
	event, err := s.task.unpackBridgeOperatorsApprovedEvent(&bridgeModels.Task{
		Data: data,
	})

	expectedPeriod := big.NewInt(9268957)
	expectedBridgeOperators := []common.Address{
		common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		common.HexToAddress("0x4a4bc674A97737376cFE990aE2fE0d2B6E738393"),
	}
	s.Equal(expectedPeriod, event.Period)
	s.Equal(expectedBridgeOperators, event.Operators)
	s.Nil(err)
}

func (s *CommonTestSuite) TestUnpackBridgeOperatorsUpdatedEventSuccess() {
	data := "7b2261646472657373223a22307862363863626264643463376233653465333963303335626433353666633562646237653261353065222c22746f70696373223a5b22307838643764353139653831633262386463363762343466643634356664326338383035313130643961623164363433653364643638623632326264653333316666222c22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030386436656465225d2c2264617461223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303230303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030323030303030303030303030303030303030303030303030306236626335626330343130373733613366383662313533376365373439356335326533386638386230303030303030303030303030303030303030303030303034613462633637346139373733373337366366653939306165326665306432623665373338333933222c22626c6f636b4e756d626572223a223078366164222c227472616e73616374696f6e48617368223a22307839356137303665643066623631643164333939623966376235616363396132616538396361616239626639353762326338613236303862343665623332666364222c227472616e73616374696f6e496e646578223a22307830222c22626c6f636b48617368223a22307838316462623832326435353735613334623365663562366266656334353762386235623762656335356338623461356463636365613132643861646238636137222c226c6f67496e646578223a22307836222c2272656d6f766564223a66616c73657d"
	event, err := s.task.unpackBridgeOperatorSetUpdatedEvent(&bridgeModels.Task{
		Data: data,
	})

	expectedPeriod := big.NewInt(9268958)
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
	fmt.Println(sig.S)
	fmt.Println(sig.R)
	fmt.Println(sig.V)
	expected := gateway.SignatureConsumerSignature{
		V: 28,
		S: [32]byte{24, 206, 69, 136, 225, 158, 70, 249, 196, 176, 181, 246, 203, 178, 45, 86, 201, 190, 6, 226, 167, 122, 59, 190, 26, 67, 150, 89, 203, 253, 245, 173},
		R: [32]byte{142, 33, 186, 68, 48, 211, 47, 96, 20, 190, 26, 42, 45, 3, 242, 225, 163, 69, 199, 168, 201, 96, 12, 95, 195, 190, 39, 63, 118, 180, 221, 83},
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
