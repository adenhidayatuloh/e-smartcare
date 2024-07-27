package dto

type CreateUpdatePemeriksaanRequest struct {
	Email      string
	Foto       string
	Tinggi     float64 `json:"tinggi"`
	Berat      float64 `json:"berat"`
	Keterangan string  `json:"keterangan"`
}
