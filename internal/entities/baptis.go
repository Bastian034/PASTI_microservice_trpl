package entities

import (
	"gereja-services/pkg/utils/date"
	"gorm.io/gorm"
	"time"
)

type BaptisEntity struct {
	NamaLengkap   string    `json:"nama_lengkap"`
	NamaAyah      string    `json:"nama_ayah"`
	NamaIbu       string    `json:"nama_ibu"`
	TempatLahir   string    `json:"tempat_lahir"`
	TanggalLahir  time.Time `json:"tanggal_lahir"`
	JenisKelamin  string    `json:"jenis_kelamin"`
	TanggalBaptis string    `json:"tanggal_baptis"`
}

type BaptisEntityModel struct {
	Entity
	BaptisEntity
}

func (BaptisEntityModel) TableName() string {
	return "baptis"
}

func (m *BaptisEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *BaptisEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
