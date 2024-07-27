package entity

import "time"

type Siswa_pemeriksaan struct {
	Email        string `gorm:"primaryKey"`
	NIS          string
	NamaLengkap  string
	TempatLahir  string
	TanggalLahir string
	Alamat       string
	NoTelepon    string
	Kelas        string
	Agama        string
	FotoProfil   string
	Pemeriksaan  Pemeriksaan1 `gorm:"foreignKey:Email;references:Email"`
}

func (Siswa_pemeriksaan) TableName() string {
	return "siswa"
}

type Pemeriksaan1 struct {
	Email      string    `gorm:"primaryKey"`
	Waktu      time.Time `gorm:"autoCreateTime"`
	Foto       string
	Tinggi     float64
	Berat      float64
	Keterangan string
}

// TableName method sets the table name to `user`
func (Pemeriksaan1) TableName() string {
	return "pemeriksaan"
}
