package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type DepositStore struct {
	*gorm.DB
}

func NewDepositStore(db *gorm.DB) *DepositStore {
	return &DepositStore{db}
}

func (d *DepositStore) Save(deposit *models.Deposit) error {
	return d.Create(deposit).Error
}
