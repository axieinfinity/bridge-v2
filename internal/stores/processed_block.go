package stores

import "gorm.io/gorm"

type ProcessedBlockStore struct {
	db *gorm.DB
}

func NewProcessedBlockStore(db *gorm.DB) *ProcessedBlockStore {
	return &ProcessedBlockStore{db: db}
}

func (b *ProcessedBlockStore) Save(height int64) error        { return nil }
func (b *ProcessedBlockStore) GetLatestBlock() (int64, error) { return -1, nil }
