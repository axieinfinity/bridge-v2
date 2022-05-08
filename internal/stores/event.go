package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type EventStore struct {
	*gorm.DB
}

func NewEventStore(db *gorm.DB) *EventStore {
	return &EventStore{db}
}

func (d *EventStore) Save(event *models.Event) error {
	return d.Create(event).Error
}
