package dto

type CreateUpdateRiwayatTanyaJawabRequest struct {
	Email      string
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
}
