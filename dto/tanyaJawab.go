package dto

type CreateUpdateTanyaJawabRequest struct {
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
	Validator  string
}

type ChatbotSimillarityRequest struct {
	Pertanyaan string `json:"pertanyaan"`
}
