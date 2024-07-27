package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"
	PemeriksaanRepository "esmartcare/repository/pemeriksaanRepository"
	"strings"
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

	urlImageNew := ""
	oldPemeriksaan, checkEmail := s.repo.GetPemeriksaanByEmail(email)

	if checkEmail != nil {
		return nil, errs.NewBadRequest("Please add email first")
	}

	urlImage, err := pkg.UploadImagePemeriksaan("foto_pemeriksaan", oldPemeriksaan.Email, ctx)

	if err != nil {
		return nil, err
	}

	if *urlImage == "" {
		return nil, errs.NewBadRequest("Image not detected")
	}

	urlImageNew = strings.Replace(*urlImage, "-temp", "", -1)

	if oldPemeriksaan.Foto != "" {
		// Delete the old image only after the new image is uploaded successfully
		errDeleteImage := pkg.DeleteImage(oldPemeriksaan.Foto)
		if errDeleteImage != nil {
			return nil, errDeleteImage
		}
	}

	// Rename the new image from temporary to final name
	err = pkg.RenameImage(*urlImage, urlImageNew)
	if err != nil {
		return nil, errs.NewInternalServerError("Error on upload image")
	}

	NewPemeriksaan := entity.Pemeriksaan{
		Email: email,

		Foto: urlImageNew,
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
