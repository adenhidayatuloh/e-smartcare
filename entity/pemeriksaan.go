package entity

import (
	"time"
)

type Pemeriksaan struct {
	Email      string    `gorm:"primaryKey"`
	Waktu      time.Time `gorm:"autoCreateTime"`
	Foto       string
	Tinggi     float64
	Berat      float64
	Keterangan string
	Siswa      Siswa `gorm:"foreignKey:Email;references:Email"`
}

// TableName method sets the table name to `user`
func (Pemeriksaan) TableName() string {
	return "pemeriksaan"
}
