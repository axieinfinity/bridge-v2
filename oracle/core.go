package oracle

import (
	"github.com/axieinfinity/bridge-v2/contract"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/vrfkey"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/chainlink/core/services/vrf/proof"
	"github.com/smartcontractkit/chainlink/core/utils"
)

type RandomWordsResponse struct {
	Proof         contract.VRFProof
	RandomRequest contract.SLARandomRequest
	RequestHash   common.Hash
}

var (
	Uint256, _ = abi.NewType("uint256", "", nil)
	Address, _ = abi.NewType("address", "", nil)
	Bytes32, _ = abi.NewType("bytes32", "", nil)
)

var seedArguments = abi.Arguments{
	{
		Name: "blockNumber",
		Type: Uint256,
	},
	{
		Name: "callbackGasLimit",
		Type: Uint256,
	},
	{
		Name: "gasPrice",
		Type: Uint256,
	},

	{
		Name: "gasFee",
		Type: Uint256,
	},

	{
		Name: "requester",
		Type: Address,
	},

	{
		Name: "consumer",
		Type: Address,
	},
	{
		Name: "nonce",
		Type: Uint256,
	},
	{
		Name: "constantFee",
		Type: Uint256,
	},
	{
		Name: "keyHash",
		Type: Bytes32,
	},

	{
		Name: "oracleAddr",
		Type: Address,
	},
}

// Returns proof seed from random request and oracle's info (hash of public key and oracle address).
func getProofSeed(req contract.SLARandomRequest, keyHash common.Hash, oracleAddr common.Address) (*big.Int, error) {
	encodedData, err := seedArguments.Pack(
		req.BlockNumber,
		req.CallbackGasLimit,
		req.GasPrice,
		req.GasFee,
		req.Requester,
		req.Consumer,
		req.Nonce,
		req.ConstantFee,
		keyHash,
		oracleAddr,
	)

	if err != nil {
		return nil, err
	}

	return utils.MustHash(string(encodedData)).Big(), nil
}

func GenerateProofResponse(key vrfkey.KeyV2, req contract.SLARandomRequest, keyHash common.Hash, oracleAddr common.Address, reqHash common.Hash) (*RandomWordsResponse, error) {
	seed, err := getProofSeed(req, keyHash, oracleAddr)
	if err != nil {
		return nil, err
	}

	pi, err := key.GenerateProof(seed)
	if err != nil {
		return nil, err
	}

	solidityProof, err := proof.SolidityPrecalculations(&pi)
	if err != nil {
		return nil, err
	}

	solidityProof.P.Seed = seed
	x, y := secp256k1.Coordinates(solidityProof.P.PublicKey)
	gx, gy := secp256k1.Coordinates(solidityProof.P.Gamma)
	cgx, cgy := secp256k1.Coordinates(solidityProof.CGammaWitness)
	shx, shy := secp256k1.Coordinates(solidityProof.SHashWitness)
	return &RandomWordsResponse{
		Proof: contract.VRFProof{
			Pk:            [2]*big.Int{x, y},
			Gamma:         [2]*big.Int{gx, gy},
			C:             solidityProof.P.C,
			S:             solidityProof.P.S,
			Seed:          seed,
			UWitness:      solidityProof.UWitness,
			CGammaWitness: [2]*big.Int{cgx, cgy},
			SHashWitness:  [2]*big.Int{shx, shy},
			ZInv:          solidityProof.ZInv,
		},
		RandomRequest: req,
		RequestHash:   reqHash,
	}, nil
}
