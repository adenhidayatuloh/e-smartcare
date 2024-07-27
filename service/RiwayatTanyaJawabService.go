package service

import (
	"esmartcare/dto"
	"esmartcare/entity"

	RiwayatTanyaJawabRepository "esmartcare/repository/riwayatTanyaJawabRepository"
)

type RiwayatTanyaJawabService interface {
	GetRiwayatByEmail(email string) ([]entity.RiwayatTanyaJawab, error)
	CreateRiwayat(request dto.CreateUpdateRiwayatTanyaJawabRequest) (entity.RiwayatTanyaJawab, error)
	DeleteRiwayatById(id uint) error
}

type riwayatTanyaJawabService struct {
	repo RiwayatTanyaJawabRepository.RiwayatTanyaJawabRepository
}

func NewRiwayatTanyaJawabService(repo RiwayatTanyaJawabRepository.RiwayatTanyaJawabRepository) RiwayatTanyaJawabService {
	return &riwayatTanyaJawabService{repo: repo}
}

func (s *riwayatTanyaJawabService) GetRiwayatByEmail(email string) ([]entity.RiwayatTanyaJawab, error) {

	return s.repo.FindByEmail(email)
}

func (s *riwayatTanyaJawabService) CreateRiwayat(request dto.CreateUpdateRiwayatTanyaJawabRequest) (entity.RiwayatTanyaJawab, error) {
	riwayat := entity.RiwayatTanyaJawab{
		Email:      request.Email,
		Pertanyaan: request.Pertanyaan,
		Jawaban:    request.Jawaban,
	}
	return s.repo.Create(riwayat)
}

func (s *riwayatTanyaJawabService) DeleteRiwayatById(id uint) error {
	return s.repo.DeleteById(id)
}
