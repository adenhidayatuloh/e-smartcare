package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg/errs"

	RiwayatTanyaJawabRepository "esmartcare/repository/riwayatTanyaJawabRepository"
)

type RiwayatTanyaJawabService interface {
	GetRiwayatByEmail(email string) ([]entity.RiwayatTanyaJawab, error)
	GetAllDataRiwayat() ([]entity.RiwayatTanyaJawab, error)
	CreateRiwayat(request dto.CreateUpdateRiwayatTanyaJawabRequest) (entity.RiwayatTanyaJawab, error)
	DeleteRiwayatById(id uint) error
	DeleteRiwayatByEmail(email string) error
}

type riwayatTanyaJawabService struct {
	repo RiwayatTanyaJawabRepository.RiwayatTanyaJawabRepository
}

// GetAllDataRiwayat implements RiwayatTanyaJawabService.
func (s *riwayatTanyaJawabService) GetAllDataRiwayat() ([]entity.RiwayatTanyaJawab, error) {
	return s.repo.GetAllDataRiwayat()
}

// DeleteRiwayatByEmail implements RiwayatTanyaJawabService.
func (s *riwayatTanyaJawabService) DeleteRiwayatByEmail(email string) error {

	riwayat, err := s.repo.FindByEmail(email)

	if err != nil {
		return err
	}

	// Cek apakah riwayat kosong
	if len(riwayat) == 0 {
		return errs.NewNotFound("No records found for the provided email")
	}
	return s.repo.DeleteByEmail(email)
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
