package dto

type CreateUpdateTanyaJawabRequest struct {
	Pertanyaan string `json:"pertanyaan" valid:"required~pertanyaan cannot be empty"`
	Jawaban    string `json:"jawaban" valid:"required~jawaban cannot be empty"`
	Validator  string
}

type ChatbotSimillarityRequest struct {
	Pertanyaan string `json:"pertanyaan"`
}
