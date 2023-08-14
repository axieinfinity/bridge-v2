package oracle

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type NewLog struct {
	Address       common.Address `json:"address" gencodec:"required"`
	Topics        []common.Hash  `json:"topics" gencodec:"required"`
	Data          hexutil.Bytes  `json:"data" gencodec:"required"`
	BlockNumber   uint64         `json:"blockNumber"`
	TxHash        common.Hash    `json:"transactionHash" gencodec:"required"`
	TxIndex       uint           `json:"transactionIndex"`
	BlockHash     common.Hash    `json:"blockHash"`
	Index         uint           `json:"logIndex"`
	Removed       bool           `json:"removed"`
	TimeStamp     uint64         `json:"timestamp"`
	PublishedTime int64          `json:"publishedTime"`
}
