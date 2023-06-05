package entities

import (
	"gereja-services/pkg/utils/date"
	"gorm.io/gorm"
)

type PemberkatanEntity struct {
	NamaJemaat    string `json:"nama_jemaat"`
	TempatLahir   string `json:"tempat_lahir"`
	TanggalLahir  string `json:"tanggal_lahir"`
	JenisKelamin  string `json:"jenis_kelamin"`
	GolonganDarah string `json:"golongan_darah"`
	Alamat        string `json:"alamat"`
	NomorHP       string `json:"nomor_hp"`
}

type PemberkatanEntityModel struct {
	Entity
	PemberkatanEntity
}

func (PemberkatanEntityModel) TableName() string {
	return "pemberkatan"
}

func (m *PemberkatanEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *PemberkatanEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
