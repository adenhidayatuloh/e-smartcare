package entity

type Alarm struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	Email        string
	Keterangan   string
	TanggalMulai string
	Jam          string
	Pengulangan  int
	Status       string
}

func (Alarm) TableName() string {
	return "alarm"
}
