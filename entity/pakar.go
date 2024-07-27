package entity

type Pakar struct {
	Email       string `gorm:"primaryKey"`
	NamaLengkap string
	Alamat      string
	NoTelepon   string
	FotoProfil  string
	User        User `gorm:"foreignKey:Email;references:Email"`
}

// TableName method sets the table name to `user`
func (Pakar) TableName() string {
	return "pakar"
}
