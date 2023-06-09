package entities

import (
	"gereja-services/pkg/utils/date"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	ID string `json:"id" gorm:"primaryKey"`

	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}

func (e *Entity) BeforeCreate(trx *gorm.DB) (err error) {
	e.CreatedAt = *date.DateTodayLocal()
	e.ID = uuid.NewString()
	return
}

func (e *Entity) BeforeUpdate(trx *gorm.DB) (err error) {
	e.ModifiedAt = date.DateTodayLocal()
	return
}
