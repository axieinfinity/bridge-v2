package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type WithdrawalStore struct {
	db *gorm.DB
}

func NewWithdrawalStore(db *gorm.DB) *WithdrawalStore {
	return &WithdrawalStore{db: db}
}

func (w *WithdrawalStore) Save(withdraw *models.Withdrawal) error   { return nil }
func (w *WithdrawalStore) Update(withdraw *models.Withdrawal) error { return nil }
func (w *WithdrawalStore) GetWithdrawalById(withdrawalId int64) (*models.Withdrawal, error) {
	return nil, nil
}
