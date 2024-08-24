package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"

	PakarRepository "esmartcare/repository/pakarRepository"

	"github.com/gin-gonic/gin"
)

type PakarService interface {
	UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreatePakarResponse, errs.MessageErr)
	CreateOrUpdatePakar(email string, payload *dto.CreatePakarRequest) (*dto.CreatePakarResponse, errs.MessageErr)
	GetPakar(email string) (*dto.CreatePakarResponse, errs.MessageErr)
}

type pakarService struct {
	PakarRepo PakarRepository.PakarRepository
}

// GetPakar implements PakarService.
func (s *pakarService) GetPakar(email string) (*dto.CreatePakarResponse, errs.MessageErr) {
	pakar, err := s.PakarRepo.GetPakarByEmail(email)

	if err != nil {
		return nil, err
	}
	pakarResponse := &dto.CreatePakarResponse{
		Email: pakar.Email,

		NamaLengkap: pakar.NamaLengkap,

		Alamat:    pakar.Alamat,
		NoTelepon: pakar.NoTelepon,

		FotoProfil: pakar.FotoProfil,
	}

	return pakarResponse, nil
}

// UpdateProfilPhoto implements pakarService.
func (s *pakarService) UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreatePakarResponse, errs.MessageErr) {

	oldPakar, checkEmail := s.PakarRepo.GetPakarByEmail(email)

	if checkEmail != nil {
		return nil, errs.NewBadRequest("Please add email first")
	}
	urlImage, err := pkg.UploadImage("foto_profil", oldPakar.Email, ctx)
	// Di sini logic nya

	if err != nil {
		return nil, err
	}

	if *urlImage == "" {
		return nil, errs.NewBadRequest("Image not detected")
	}

	NewPakar := entity.Pakar{
		Email: email,

		FotoProfil: *urlImage,
	}

	// Update the student record
	updatedUser, err := s.PakarRepo.UpdatePakar(oldPakar, &NewPakar)
	if err != nil {
		return nil, errs.NewBadRequest("Cannot update Pakar")
	}

	updatePakarResponse := &dto.CreatePakarResponse{
		Email: updatedUser.Email,

		NamaLengkap: updatedUser.NamaLengkap,

		Alamat:    updatedUser.Alamat,
		NoTelepon: updatedUser.NoTelepon,

		FotoProfil: updatedUser.FotoProfil,
	}

	return updatePakarResponse, nil

}

func NewpakarService(pakarRepo PakarRepository.PakarRepository) PakarService {
	return &pakarService{pakarRepo}
}

// CreateOrUpdatePakar implements pakarService.
func (s *pakarService) CreateOrUpdatePakar(email string, payload *dto.CreatePakarRequest) (*dto.CreatePakarResponse, errs.MessageErr) {

	NewPakar := entity.Pakar{
		Email: email,

		NamaLengkap: payload.NamaLengkap,

		Alamat:    payload.Alamat,
		NoTelepon: payload.NoTelepon,
	}

	oldPakar, checkEmail := s.PakarRepo.GetPakarByEmail(email)

	if checkEmail == nil {

		updatedUser, err := s.PakarRepo.UpdatePakar(oldPakar, &NewPakar)
		if err != nil {
			return nil, err
		}
		updatePakarResponse := &dto.CreatePakarResponse{
			Email: updatedUser.Email,

			NamaLengkap: updatedUser.NamaLengkap,

			Alamat:    updatedUser.Alamat,
			NoTelepon: updatedUser.NoTelepon,

			FotoProfil: updatedUser.FotoProfil,
		}

		return updatePakarResponse, nil
	}

	// Create the new student record
	CreatedUser, err := s.PakarRepo.CreatePakar(&NewPakar)
	if err != nil {
		return nil, err
	}

	CreatePakarResponse := &dto.CreatePakarResponse{
		Email: CreatedUser.Email,

		NamaLengkap: CreatedUser.NamaLengkap,

		Alamat:    CreatedUser.Alamat,
		NoTelepon: CreatedUser.NoTelepon,

		FotoProfil: CreatedUser.FotoProfil,
	}

	return CreatePakarResponse, nil
}
