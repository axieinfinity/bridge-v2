package task

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
)

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
