package oracle

import (
	"encoding/hex"
	"log"

	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/vrfkey"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
)

func GenerateVRFKey() {
	key, err := vrfkey.NewV2()
	if err != nil {
		panic(err)
	}
	point, err := key.PublicKey.Point()
	x, y := secp256k1.Coordinates(point)

	log.Printf("Generated public key is: %s (x=%v,y=%v)", key.PublicKey, x, y)
	log.Printf("Key hash is: %v", key.PublicKey.MustHash())
	log.Printf("Secret key is: 0x%v", hex.EncodeToString(key.Raw()))
}
