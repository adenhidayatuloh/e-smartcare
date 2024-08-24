package dto

type CreateUpdateTanyaJawabRequest struct {
	Pertanyaan string `json:"pertanyaan" valid:"required~Pertanyaan cannot be empty "`
	Jawaban    string `json:"jawaban" valid:"required~Jawaban cannot be empty"`
	Validator  string
}

type ChatbotSimillarityRequest struct {
	Pertanyaan string `json:"pertanyaan"`
}
