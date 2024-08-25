package entity

type Siswa struct {
	Email        string `gorm:"primaryKey"`
	NIS          string
	NamaLengkap  string
	TempatLahir  string
	TanggalLahir string `gorm:"default:NULL"`
	Alamat       string
	NoTelepon    string
	Kelas        string
	Agama        string
	FotoProfil   string
	User         User `gorm:"foreignKey:Email;references:Email"`
}

// TableName method sets the table name to `user`
func (Siswa) TableName() string {
	return "siswa"
}
