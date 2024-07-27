package entity

import (
	"time"
)

type RiwayatTanyaJawab struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	Email      string
	Waktu      time.Time `gorm:"autoCreateTime"`
	Pertanyaan string
	Jawaban    string
}

// TableName method sets the table name to `user`
func (RiwayatTanyaJawab) TableName() string {
	return "riwayat_tanya_jawab"
}
