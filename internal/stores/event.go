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

func (e *EventStore) Save(event *models.Event) error {
	return e.Create(event).Error
}

func (e *EventStore) DeleteEvents(createdAt uint64) error {
	return e.Where("created_at <= ?", createdAt).Delete(&models.Event{}).Error
}

func (e *EventStore) Count() int64 {
	var count int64
	e.Model(&models.Event{}).Select("id").Count(&count)
	return count
}
