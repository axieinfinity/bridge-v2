package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type WithdrawalStore struct {
	*gorm.DB
}

func NewWithdrawalStore(db *gorm.DB) *WithdrawalStore {
	return &WithdrawalStore{db}
}

func (w *WithdrawalStore) Save(withdraw *models.Withdrawal) error {
	return w.Create(withdraw).Error
}

func (w *WithdrawalStore) Update(withdraw *models.Withdrawal) error {
	return w.Updates(withdraw).Error
}

func (w *WithdrawalStore) GetWithdrawalById(withdrawalId int64) (*models.Withdrawal, error) {
	var withdraw *models.Withdrawal
	err := w.Model(&models.Withdrawal{}).Where("withdrawal_id = ?", withdrawalId).First(withdraw).Error
	return withdraw, err
}
