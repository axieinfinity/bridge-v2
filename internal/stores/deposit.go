package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type depositStore struct {
	*gorm.DB
}

func NewDepositStore(db *gorm.DB) *depositStore {
	return &depositStore{db}
}

func (d *depositStore) Save(deposit *models.Deposit) error {
	return d.Create(deposit).Error
}
