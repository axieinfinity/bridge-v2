package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProcessedReceiptStore struct {
	*gorm.DB
}

func NewProcessedReceiptStore(db *gorm.DB) *ProcessedReceiptStore {
	return &ProcessedReceiptStore{db}
}

func (b *ProcessedReceiptStore) Save(taskId int, receiptId int64) error {
	return b.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.ProcessedReceipt{TaskId: taskId, ReceiptId: receiptId}).Error
}
