package task

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/vrfkey"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"strings"
)

var VRFConfig *VRF

type VRF struct {
	WaitForBlock    int          `json:"waitForBlock"`
	ContractAddress string       `json:"contractAddress"`
	ContractName    string       `json:"contractName"`
	SecretKey       string       `json:"secretKey"`
	Key             vrfkey.KeyV2 `json:"-"`
	KeyHash         common.Hash  `json:"-"`
}

type BridgeOperatorsSorter []common.Address

func (b BridgeOperatorsSorter) Len() int { return len(b) }

func (b BridgeOperatorsSorter) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b BridgeOperatorsSorter) Less(i, j int) bool {
	return bytes.Compare(b[i].Bytes(), b[j].Bytes()) == -1
}

func EqualOperatorSet(a, b []common.Address) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if bytes.Compare(v.Bytes(), b[i].Bytes()) != 0 {
			return false
		}
	}

	return true
}

func (v *VRF) SetVRFKey() {
	if strings.HasPrefix(v.SecretKey, "0x") {
		v.SecretKey = v.SecretKey[2:]
	}
	byteSK, err := hex.DecodeString(v.SecretKey)
	if err != nil {
		panic(err)
	}
	v.Key = vrfkey.Raw(byteSK).Key()
	publicKey, err := secp256k1.NewPublicKeyFromBytes(v.Key.PublicKey[:])
	if err != nil {
		panic(err)
	}
	v.KeyHash = publicKey.MustHash()
}
