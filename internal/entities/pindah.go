package entities

import (
	"gereja-services/pkg/utils/date"
	"gorm.io/gorm"
)

type PindahEntity struct {
	NamaJemaat    string `json:"nama_jemaat"`
	TempatLahir   string `json:"tempat_lahir"`
	TanggalLahir  string `json:"tanggal_lahir"`
	JenisKelamin  string `json:"jenis_kelamin"`
	GolonganDarah string `json:"golongan_darah"`
	Alamat        string `json:"alamat"`
	NomorHP       string `json:"nomor_hp"`
	Gereja        string `json:"gereja"`
}

type PindahEntityModel struct {
	Entity
	PemberkatanEntity
}

func (PindahEntityModel) TableName() string {
	return "pindah"
}

func (m *PindahEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *PindahEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
