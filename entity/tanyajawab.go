package entity

type TanyaJawab struct {
	IDTanyaJawab int    `gorm:"primaryKey;autoIncrement"`
	Pertanyaan   string `json:"pertanyaan"`
	Jawaban      string `json:"jawaban"`
	Validator    string `json:"validator"`
}

// TableName method sets the table name to `user`
func (TanyaJawab) TableName() string {
	return "tanya_jawab"
}

type FAQ struct {
	Question string
	Answer   string
	Topic    string
}
