package entity

type Admin struct {
	Email       string `gorm:"primaryKey"`
	NamaLengkap string
	Alamat      string
	NoTelepon   string
	FotoProfil  string
	User        User `gorm:"foreignKey:Email;references:Email"`
}

// TableName method sets the table name to `user`
func (Admin) TableName() string {
	return "admin"
}
