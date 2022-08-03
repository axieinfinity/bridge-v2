package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type processedReceiptStore struct {
	*gorm.DB
}

func NewProcessedReceiptStore(db *gorm.DB) ProcessedReceiptStore {
	return &processedReceiptStore{db}
}

func (b *processedReceiptStore) Save(taskId int, receiptId int64) error {
	return b.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.ProcessedReceipt{TaskId: taskId, ReceiptId: receiptId}).Error
}
