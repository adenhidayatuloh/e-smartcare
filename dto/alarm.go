package dto

type CreateUpdateAlarmRequestResponse struct {
	Email        string
	Keterangan   string `json:"keterangan"`
	TanggalMulai string `json:"tanggal_mulai" valid:"required~Tanggal mulai cannot be empty"`
	Jam          string `json:"jam" valid:"required~Jam cannot be empty" `
	Pengulangan  int    `json:"pengulangan" valid:"required~Pengulangan cannot be empty"`
	Status       string `json:"status" valid:"required~Status alarm cannot be empty,statusAlarm~Status hanya bisa diisi 0 atau 1"`
}
