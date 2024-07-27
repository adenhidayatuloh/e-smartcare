package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"
	"strings"

	PakarRepository "esmartcare/repository/pakarRepository"

	"github.com/gin-gonic/gin"
)

type PakarService interface {
	CreateOrUpdatePakarWithProfilPhoto(email string, payload *dto.CreatePakarRequest, ctx *gin.Context) (*dto.CreatePakarResponse, errs.MessageErr)
	UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreatePakarResponse, errs.MessageErr)
	CreateOrUpdatePakar(email string, payload *dto.CreatePakarRequest) (*dto.CreatePakarResponse, errs.MessageErr)

	//Update(email string, payload *dto.CreatePakarRequestAndResponse) (*dto.CreatePakarRequestAndResponse, errs.MessageErr)
	//UpdatePhoto(email string) (*dto.UpdatePhotoResponse, errs.MessageErr)
	// Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	// GetAllUsers(jenis_akun string) ([]dto.GetAllUsersResponse, errs.MessageErr)
	// GetAllUsersNotValidate(jenis_akun string) ([]dto.GetAllUsersResponse, errs.MessageErr)
	// UpdateUser(email string) (*dto.UpdateUserResponse, errs.MessageErr)
	//DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr)
}

type pakarService struct {
	PakarRepo PakarRepository.PakarRepository
}

// UpdateProfilPhoto implements pakarService.
func (s *pakarService) UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreatePakarResponse, errs.MessageErr) {

	urlImageNew := ""
	oldPakar, checkEmail := s.PakarRepo.GetPakarByEmail(email)

	if checkEmail != nil {
		return nil, errs.NewBadRequest("Please add email first")
	}

	urlImage, err := pkg.UploadImage("foto_profil", oldPakar.Email, ctx)

	if err != nil {
		return nil, err
	}

	if *urlImage == "" {
		return nil, errs.NewBadRequest("Image not detected")
	}

	urlImageNew = strings.Replace(*urlImage, "-temp", "", -1)

	if oldPakar.FotoProfil != "" {
		// Delete the old image only after the new image is uploaded successfully
		errDeleteImage := pkg.DeleteImage(oldPakar.FotoProfil)
		if errDeleteImage != nil {
			return nil, errDeleteImage
		}
	}

	// Rename the new image from temporary to final name
	err = pkg.RenameImage(*urlImage, urlImageNew)
	if err != nil {
		return nil, errs.NewInternalServerError("Error on upload image")
	}

	NewPakar := entity.Pakar{
		Email: email,

		FotoProfil: urlImageNew,
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

func (s *pakarService) CreateOrUpdatePakarWithProfilPhoto(email string, payload *dto.CreatePakarRequest, ctx *gin.Context) (*dto.CreatePakarResponse, errs.MessageErr) {

	urlImageNew := ""
	NewPakar := entity.Pakar{
		Email: email,

		NamaLengkap: payload.NamaLengkap,

		Alamat:    payload.Alamat,
		NoTelepon: payload.NoTelepon,

		FotoProfil: payload.FotoProfil,
	}

	oldPakar, checkEmail := s.PakarRepo.GetPakarByEmail(email)

	// Upload image with a temporary name
	urlImage, err := pkg.UploadImage("foto_profil", oldPakar.Email, ctx)

	if err != nil {
		return nil, errs.NewBadRequest("Cannot upload foto profil Pakar")
	}

	if *urlImage != "" {
		urlImageNew = strings.Replace(*urlImage, "-temp", "", -1)
	}

	if checkEmail == nil {
		if oldPakar.FotoProfil != "" && *urlImage != "" {
			// Delete the old image only after the new image is uploaded successfully
			errDeleteImage := pkg.DeleteImage(oldPakar.FotoProfil)
			if errDeleteImage != nil {
				return nil, errDeleteImage
			}
		}

		if *urlImage != "" {
			// Rename the new image from temporary to final name
			err := pkg.RenameImage(*urlImage, urlImageNew)
			if err != nil {
				return nil, errs.NewInternalServerError("Cannot rename image")
			}
			NewPakar.FotoProfil = urlImageNew
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

	if *urlImage != "" {
		// Rename the new image from temporary to final name
		err := pkg.RenameImage(*urlImage, urlImageNew)
		if err != nil {
			return nil, errs.NewInternalServerError("Cannot rename image")
		}
		NewPakar.FotoProfil = urlImageNew
	}

	// Create the new student record
	CreatedUser, err := s.PakarRepo.CreatePakar(&NewPakar)
	if err != nil {
		return nil, err
	}

	CreatePakarResponse := &dto.CreatePakarResponse{
		Email: CreatedUser.Email,

		NamaLengkap: CreatedUser.NamaLengkap,

		Alamat:     CreatedUser.Alamat,
		NoTelepon:  CreatedUser.NoTelepon,
		FotoProfil: CreatedUser.FotoProfil,
	}

	return CreatePakarResponse, nil
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
