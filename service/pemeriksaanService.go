package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"
	PemeriksaanRepository "esmartcare/repository/pemeriksaanRepository"
	"time"

	"github.com/gin-gonic/gin"
)

type PemeriksaanService interface {
	GetAllPemeriksaan() ([]entity.Pemeriksaan, error)
	CreatePemeriksaan(request dto.CreateUpdatePemeriksaanRequest) (*entity.Pemeriksaan, error)
	GetPemeriksaanByEmail(email string) ([]entity.Pemeriksaan, error)
	DeletePemeriksaanByEmail(email string) error
	UpdatePhotoPemeriksaan(email string, ctx *gin.Context) (*dto.CreateUpdatePemeriksaanRequest, errs.MessageErr)
}

type pemeriksaanService struct {
	repo PemeriksaanRepository.PemeriksaanRepository
}

func NewPemeriksaanService(repo PemeriksaanRepository.PemeriksaanRepository) PemeriksaanService {
	return &pemeriksaanService{repo: repo}
}

func (s *pemeriksaanService) GetAllPemeriksaan() ([]entity.Pemeriksaan, error) {
	return s.repo.FindAll()
}

func (s *pemeriksaanService) CreatePemeriksaan(request dto.CreateUpdatePemeriksaanRequest) (*entity.Pemeriksaan, error) {
	pemeriksaan := entity.Pemeriksaan{
		Email:      request.Email,
		Tinggi:     request.Tinggi,
		Berat:      request.Berat,
		Keterangan: request.Keterangan,
	}

	oldPemeriksaan, checkEmail := s.repo.GetPemeriksaanByEmail(pemeriksaan.Email)

	if checkEmail == nil {

		pemeriksaan.Waktu = time.Now()

		updatedUser, err := s.repo.UpdatePemeriksaan(oldPemeriksaan, &pemeriksaan)
		if err != nil {
			return nil, err
		}

		return updatedUser, nil
	}

	return s.repo.Create(pemeriksaan)
}

func (s *pemeriksaanService) GetPemeriksaanByEmail(email string) ([]entity.Pemeriksaan, error) {
	return s.repo.FindByEmail(email)
}

func (s *pemeriksaanService) DeletePemeriksaanByEmail(email string) error {
	return s.repo.DeleteByEmail(email)
}

// UpdateProfilPhoto implements PemeriksaanService.
func (s *pemeriksaanService) UpdatePhotoPemeriksaan(email string, ctx *gin.Context) (*dto.CreateUpdatePemeriksaanRequest, errs.MessageErr) {

	oldPemeriksaan, checkEmail := s.repo.GetPemeriksaanByEmail(email)

	if checkEmail != nil {
		return nil, errs.NewBadRequest("Please add email first")
	}

	newKeyImage := oldPemeriksaan.Email + "-pemeriksaan"

	urlImage, err := pkg.UploadImage("foto_pemeriksaan", newKeyImage, ctx)
	// Di sini logic nya

	if err != nil {
		return nil, err
	}

	if *urlImage == "" {
		return nil, errs.NewBadRequest("Image not detected")
	}
	NewPemeriksaan := entity.Pemeriksaan{
		Email: email,

		Foto: *urlImage,
	}

	// Update the student record
	updatedUser, err := s.repo.UpdatePemeriksaan(oldPemeriksaan, &NewPemeriksaan)
	if err != nil {
		return nil, errs.NewBadRequest("Cannot update Pemeriksaan")
	}

	updatePemeriksaanResponse := &dto.CreateUpdatePemeriksaanRequest{
		Email: updatedUser.Email,
		Foto:  NewPemeriksaan.Foto,
	}

	return updatePemeriksaanResponse, nil

}
