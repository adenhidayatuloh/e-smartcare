package entity

import (
	"time"
)

type Pemeriksaan struct {
	IdPemeriksaan int `gorm:"primaryKey"`
	Email         string
	Waktu         time.Time `gorm:"autoCreateTime"`
	Foto          string
	Tinggi        float64
	Berat         float64
	Keterangan    string
}

// TableName method sets the table name to `user`
func (Pemeriksaan) TableName() string {
	return "pemeriksaan"
}
