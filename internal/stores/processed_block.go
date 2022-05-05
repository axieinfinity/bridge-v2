package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type ProcessedBlockStore struct {
	*gorm.DB
}

func NewProcessedBlockStore(db *gorm.DB) *ProcessedBlockStore {
	return &ProcessedBlockStore{db}
}

func (b *ProcessedBlockStore) Save(chainId string, height int64) error {
	return b.Create(&models.ProcessedBlock{Block: height, ChainId: chainId}).Error
}

func (b *ProcessedBlockStore) GetLatestBlock(chainId string) (int64, error) {
	var block *models.ProcessedBlock
	if err := b.Model(&models.ProcessedBlock{}).Where("chain_id = ?", chainId).Last(&block).Error; err != nil {
		return -1, err
	}
	return block.Block, nil
}
