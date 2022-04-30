package listener

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type EthBlock struct {
	block     *types.Block
	txs       []ITransaction
	receipts  []IReceipt
}

func NewEthBlock(rpcUrl string, block *types.Block) (*EthBlock, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ethBlock := &EthBlock{
		block:    block,
		txs:      make([]ITransaction, len(block.Transactions())),
		receipts: make([]IReceipt, len(block.Transactions())),
	}
	for i, tx := range block.Transactions() {
		// get receipt from tx Hash
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return nil, err
		}
		ethTx, err := NewEthTransaction(rpcUrl, tx)
		if err != nil {
			return nil, err
		}
		ethReceipt := &EthReceipt{
			receipt: receipt,
			tx:      ethTx,
		}
		ethBlock.txs[i] = ethTx
		ethBlock.receipts[i] = ethReceipt
	}
	return ethBlock, nil
}

func (b *EthBlock) GetHash() common.Hash { return b.block.Hash() }
func (b *EthBlock) GetHeight() uint64 { return b.block.NumberU64() }

func (b *EthBlock) GetTransactions() []ITransaction {
	return b.txs
}

func (b *EthBlock) GetReceipts() []IReceipt {
	if len(b.receipts) > 0 {
		return b.receipts
	}
	for _, tx := range b.GetTransactions() {
		b.receipts = append(b.receipts, tx.GetReceipt())
	}
	return b.receipts
}

type EthTransaction struct {
	chainId   *big.Int
	sender    common.Address
	tx        *types.Transaction
	receipt   IReceipt
}

func NewEthTransaction(rpcUrl string, tx *types.Transaction) (*EthTransaction, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	sender, err := types.LatestSignerForChainID(chainId).Sender(tx)
	if err != nil {
		return nil, err
	}
	ethTx := &EthTransaction{
		chainId: chainId,
		sender:  sender,
		tx:      tx,
	}
	r, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, err
	}
	ethTx.receipt = &EthReceipt{
		receipt: r,
		tx: ethTx,
	}
	return ethTx, nil
}

func (b *EthTransaction) GetHash() common.Hash {
	return b.tx.Hash()
}

func (b *EthTransaction) GetFromAddress() string {
	return b.sender.Hex()
}
func (b *EthTransaction) GetToAddress() string {
	if b.tx.To() != nil {
		return b.tx.To().Hex()
	}
	return ""
}

func (b *EthTransaction) GetData() []byte {
	return b.tx.Data()
}

func (b *EthTransaction) GetValue() *big.Int {
	return b.tx.Value()
}

func (b *EthTransaction) GetStatus() bool {
	return b.receipt.GetStatus()
}

func (b *EthTransaction) GetReceipt() IReceipt {
	return b.receipt
}

type EthReceipt struct {
	receipt *types.Receipt
	tx      *EthTransaction
}

func (r *EthReceipt) GetTransaction() ITransaction {
	return r.tx
}

func (r *EthReceipt) GetStatus() bool {
	return r.receipt.Status == 1
}

func (r *EthReceipt) GetLogs() (logs []ILog) {
	for _, l := range r.receipt.Logs {
		logs = append(logs, &EthLog{l})
	}
	return
}

type EthLog struct {
	l *types.Log
}

func(e *EthLog) GetContractAddress() string {
	return e.l.Address.Hex()
}

func(e *EthLog) GetTopics() (topics []string) {
	for _, topic := range e.l.Topics {
		topics = append(topics, topic.Hex())
	}
	return
}

func(e *EthLog) GetData() []byte {
	return e.l.Data
}

func(e *EthLog) GetIndex() uint {
	return e.l.Index
}
