package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"fmt"
	"log"
	"strconv"

	"github.com/blevesearch/bleve/v2"

	TanyaJawabRepository "esmartcare/repository/tanyaJawabRepository"
)

type TanyaJawabService interface {
	GetAllTanyaJawab() ([]entity.TanyaJawab, error)
	GetTanyaJawabByValidationStatus(isValidated bool) ([]entity.TanyaJawab, error)
	GetTanyaJawabByID(id int) (entity.TanyaJawab, error)
	CreateTanyaJawab(request dto.CreateUpdateTanyaJawabRequest) (entity.TanyaJawab, error)
	UpdateTanyaJawab(id int, request dto.CreateUpdateTanyaJawabRequest) (entity.TanyaJawab, error)
	UpdateValidator(id int, validator string) (entity.TanyaJawab, error)
	DeleteTanyaJawab(id int) error
	GetSimillaryQuestion(newQuestion string) ([]string, error)
	GetChatQuestion(newQuestion string) (string, string, error)
	GetChatBotUpdate() ([]entity.FAQ, error)
}

type tanyaJawabService struct {
	repo TanyaJawabRepository.TanyaJawabRepository
}

// GetChatBotUpdate implements TanyaJawabService.
func (s *tanyaJawabService) GetChatBotUpdate() ([]entity.FAQ, error) {
	return s.repo.FindForChatbot()
}

func NewTanyaJawabService(repo TanyaJawabRepository.TanyaJawabRepository) TanyaJawabService {
	return &tanyaJawabService{repo: repo}
}

func (s *tanyaJawabService) GetAllTanyaJawab() ([]entity.TanyaJawab, error) {
	return s.repo.FindAll()
}

func (s *tanyaJawabService) GetTanyaJawabByValidationStatus(isValidated bool) ([]entity.TanyaJawab, error) {
	return s.repo.FindByValidationStatus(isValidated)
}

func (s *tanyaJawabService) GetTanyaJawabByID(id int) (entity.TanyaJawab, error) {
	return s.repo.FindByID(id)
}

func (s *tanyaJawabService) CreateTanyaJawab(request dto.CreateUpdateTanyaJawabRequest) (entity.TanyaJawab, error) {
	tanyaJawab := entity.TanyaJawab{
		Pertanyaan: request.Pertanyaan,
		Jawaban:    request.Jawaban,
	}
	return s.repo.Create(tanyaJawab)
}

func (s *tanyaJawabService) UpdateTanyaJawab(id int, request dto.CreateUpdateTanyaJawabRequest) (entity.TanyaJawab, error) {
	tanyaJawab, err := s.repo.FindByID(id)
	if err != nil {
		return entity.TanyaJawab{}, err
	}

	tanyaJawab.Pertanyaan = request.Pertanyaan
	tanyaJawab.Jawaban = request.Jawaban

	return s.repo.Update(tanyaJawab)
}

func (s *tanyaJawabService) UpdateValidator(id int, validator string) (entity.TanyaJawab, error) {
	tanyaJawab, err := s.repo.FindByID(id)
	if err != nil {
		return entity.TanyaJawab{}, err
	}

	tanyaJawab.Validator = validator

	return s.repo.Update(tanyaJawab)
}

func (s *tanyaJawabService) DeleteTanyaJawab(id int) error {

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *tanyaJawabService) GetSimillaryQuestion(newQuestion string) ([]string, error) {
	type FAQ struct {
		Question string
		Answer   string
		Topic    string
	}

	var faqs []FAQ
	listSimilarity := []string{}

	data, err := s.repo.FindByValidationStatus(true)

	for _, v := range data {

		newFaqs := FAQ{
			Question: v.Pertanyaan,
			Answer:   v.Jawaban,
		}

		faqs = append(faqs, newFaqs)
	}

	if err != nil {
		return nil, err
	}

	// Buat atau buka indeks Bleve
	indexMapping := bleve.NewIndexMapping()
	index, err := bleve.Open("faq.bleve")

	if err == bleve.ErrorIndexPathDoesNotExist {

		faqsIndex, err := s.GetChatBotUpdate()

		index, err = pkg.CreateIndex("faq.bleve", indexMapping)
		if err != nil {
			log.Fatal(err)
		}

		// Hapus semua dokumen dari indeks sebelum pengindeksan ulang
		if err := pkg.DeleteAllDocuments(index); err != nil {
			log.Fatal(err)
		}

		// Indexing data
		for i, faq := range faqsIndex {
			err := index.Index(fmt.Sprintf("%d", i), faq)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else if err != nil {
		return nil, err
	}

	defer index.Close()

	questionQuery := bleve.NewMatchQuery(newQuestion)
	questionQuery.SetField("Question")

	// Query untuk mencocokkan topik
	topicQuery := bleve.NewMatchQuery(newQuestion)
	topicQuery.SetField("Topic")

	mainQuery := bleve.NewBooleanQuery()
	mainQuery.AddMust(questionQuery) // Pertanyaan baru harus cocok
	mainQuery.AddMust(topicQuery)

	searchRequest := bleve.NewSearchRequest(mainQuery)
	searchRequest.Size = 5 // Batasi jumlah hasil yang ditampilkan
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Menampilkan daftar pertanyaan yang paling relevan
	fmt.Println("Pertanyaan:", newQuestion)
	fmt.Println("Rekomendasi Pertanyaan:")
	for _, hit := range searchResult.Hits {
		index, _ := strconv.Atoi(hit.ID)
		similarity := hit.Score * 100
		rekomendasi := fmt.Sprintf("%s (%.2f%% kemiripan)", faqs[index].Question, similarity)

		listSimilarity = append(listSimilarity, rekomendasi)

	}

	return listSimilarity, nil

}

func (s *tanyaJawabService) GetChatQuestion(newQuestion string) (string, string, error) {
	type FAQ struct {
		Question string
		Answer   string
		Topic    string
	}

	var faqs []FAQ

	data, err := s.repo.FindByValidationStatus(true)

	for _, v := range data {

		newFaqs := FAQ{
			Question: v.Pertanyaan,
			Answer:   v.Jawaban,
			Topic:    "stunting",
		}

		faqs = append(faqs, newFaqs)
	}

	if err != nil {
		return "", "", err
	}

	// Buat atau buka indeks Bleve
	indexMapping := bleve.NewIndexMapping()
	index, err := bleve.Open("faq.bleve")

	if err == bleve.ErrorIndexPathDoesNotExist {

		faqsIndex, err := s.GetChatBotUpdate()

		index, err = pkg.CreateIndex("faq.bleve", indexMapping)
		if err != nil {
			log.Fatal(err)
		}

		// Hapus semua dokumen dari indeks sebelum pengindeksan ulang
		if err := pkg.DeleteAllDocuments(index); err != nil {
			log.Fatal(err)
		}

		// Indexing data
		for i, faq := range faqsIndex {
			err := index.Index(fmt.Sprintf("%d", i), faq)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else if err != nil {
		return "", "", err
	}

	defer index.Close()

	// Pencarian di indeks untuk mendapatkan daftar pertanyaan relevan
	questionQuery := bleve.NewMatchQuery(newQuestion)
	questionQuery.SetField("Question")

	// Query untuk mencocokkan topik
	topicQuery := bleve.NewMatchQuery(newQuestion)
	topicQuery.SetField("Topic")

	mainQuery := bleve.NewBooleanQuery()
	mainQuery.AddMust(questionQuery) // Pertanyaan baru harus cocok
	mainQuery.AddMust(topicQuery)
	searchRequest := bleve.NewSearchRequest(mainQuery)
	searchRequest.Size = 5 // Batasi jumlah hasil yang ditampilkan
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	jawaban := ""
	kemiripan := ""

	if len(searchResult.Hits) > 0 {
		mostRelevantIndex := searchResult.Hits[0].ID

		mostRelevantIndexInt, _ := strconv.Atoi(mostRelevantIndex)

		fmt.Println(mostRelevantIndexInt)

		fmt.Println(len(faqs))
		jawaban = faqs[mostRelevantIndexInt].Answer
		similarity := searchResult.Hits[0].Score * 100
		kemiripan = fmt.Sprintf("%.2f%%", similarity)
	} else {
		jawaban = "Jawaban tidak ditemukan"
	}

	return jawaban, kemiripan, nil
}
