package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"
	"strings"

	AdminRepository "esmartcare/repository/adminRepository"

	"github.com/gin-gonic/gin"
)

type AdminService interface {
	CreateOrUpdateAdminWithProfilPhoto(email string, payload *dto.CreateAdminRequest, ctx *gin.Context) (*dto.CreateAdminResponse, errs.MessageErr)
	UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreateAdminResponse, errs.MessageErr)
	CreateOrUpdateAdmin(email string, payload *dto.CreateAdminRequest) (*dto.CreateAdminResponse, errs.MessageErr)
	GetAdmin(email string) (*dto.CreateAdminResponse, errs.MessageErr)
}

type adminService struct {
	AdminRepo AdminRepository.AdminRepository
}

// GetAdmin implements AdminService.
func (s *adminService) GetAdmin(email string) (*dto.CreateAdminResponse, errs.MessageErr) {
	admin, err := s.AdminRepo.GetAdminByEmail(email)

	if err != nil {
		return nil, err
	}
	AdminResponse := &dto.CreateAdminResponse{
		Email: admin.Email,

		NamaLengkap: admin.NamaLengkap,

		Alamat:    admin.Alamat,
		NoTelepon: admin.NoTelepon,

		FotoProfil: admin.FotoProfil,
	}

	return AdminResponse, nil

}

// UpdateProfilPhoto implements adminService.
func (s *adminService) UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreateAdminResponse, errs.MessageErr) {
	oldAdmin, checkEmail := s.AdminRepo.GetAdminByEmail(email)

	if checkEmail != nil {
		return nil, errs.NewBadRequest("Please add email first")
	}

	urlImage, err := pkg.UploadImage("foto_profil", oldAdmin.Email, ctx)
	// Di sini logic nya

	if err != nil {
		return nil, err
	}

	if *urlImage == "" {
		return nil, errs.NewBadRequest("Image not detected")
	}

	NewAdmin := entity.Admin{
		Email: email,

		FotoProfil: *urlImage,
	}

	// Update the student record
	updatedUser, err := s.AdminRepo.UpdateAdmin(oldAdmin, &NewAdmin)
	if err != nil {
		return nil, errs.NewBadRequest("Cannot update Admin")
	}

	updateAdminResponse := &dto.CreateAdminResponse{
		Email: updatedUser.Email,

		NamaLengkap: updatedUser.NamaLengkap,

		Alamat:    updatedUser.Alamat,
		NoTelepon: updatedUser.NoTelepon,

		FotoProfil: updatedUser.FotoProfil,
	}

	return updateAdminResponse, nil

}

func NewadminService(adminRepo AdminRepository.AdminRepository) AdminService {
	return &adminService{AdminRepo: adminRepo}
}

func (s *adminService) CreateOrUpdateAdminWithProfilPhoto(email string, payload *dto.CreateAdminRequest, ctx *gin.Context) (*dto.CreateAdminResponse, errs.MessageErr) {

	urlImageNew := ""
	NewAdmin := entity.Admin{
		Email: email,

		NamaLengkap: payload.NamaLengkap,

		Alamat:    payload.Alamat,
		NoTelepon: payload.NoTelepon,

		FotoProfil: payload.FotoProfil,
	}

	oldAdmin, checkEmail := s.AdminRepo.GetAdminByEmail(email)

	// Upload image with a temporary name
	urlImage, err := pkg.UploadImage("foto_profil", oldAdmin.Email, ctx)

	if err != nil {
		return nil, errs.NewBadRequest("Cannot upload foto profil Admin")
	}

	if *urlImage != "" {
		urlImageNew = strings.Replace(*urlImage, "-temp", "", -1)
	}

	if checkEmail == nil {
		if oldAdmin.FotoProfil != "" && *urlImage != "" {
			// Delete the old image only after the new image is uploaded successfully
			errDeleteImage := pkg.DeleteImage(oldAdmin.FotoProfil)
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
			NewAdmin.FotoProfil = urlImageNew
		}

		// Update the student record
		updatedUser, err := s.AdminRepo.UpdateAdmin(oldAdmin, &NewAdmin)
		if err != nil {
			return nil, errs.NewBadRequest("Cannot update Admin")
		}
		updateAdminResponse := &dto.CreateAdminResponse{
			Email: updatedUser.Email,

			NamaLengkap: updatedUser.NamaLengkap,

			Alamat:    updatedUser.Alamat,
			NoTelepon: updatedUser.NoTelepon,

			FotoProfil: updatedUser.FotoProfil,
		}

		return updateAdminResponse, nil
	}

	if *urlImage != "" {
		// Rename the new image from temporary to final name
		err := pkg.RenameImage(*urlImage, urlImageNew)
		if err != nil {
			return nil, errs.NewInternalServerError("Cannot rename image")
		}
		NewAdmin.FotoProfil = urlImageNew
	}

	// Create the new student record
	CreatedUser, err := s.AdminRepo.CreateAdmin(&NewAdmin)
	if err != nil {
		return nil, err
	}

	CreateAdminResponse := &dto.CreateAdminResponse{
		Email: CreatedUser.Email,

		NamaLengkap: CreatedUser.NamaLengkap,

		Alamat:     CreatedUser.Alamat,
		NoTelepon:  CreatedUser.NoTelepon,
		FotoProfil: CreatedUser.FotoProfil,
	}

	return CreateAdminResponse, nil
}

// CreateOrUpdateAdmin implements adminService.
func (s *adminService) CreateOrUpdateAdmin(email string, payload *dto.CreateAdminRequest) (*dto.CreateAdminResponse, errs.MessageErr) {

	NewAdmin := entity.Admin{
		Email: email,

		NamaLengkap: payload.NamaLengkap,

		Alamat:    payload.Alamat,
		NoTelepon: payload.NoTelepon,
	}

	oldAdmin, checkEmail := s.AdminRepo.GetAdminByEmail(email)

	if checkEmail == nil {

		updatedUser, err := s.AdminRepo.UpdateAdmin(oldAdmin, &NewAdmin)
		if err != nil {
			return nil, err
		}
		updateAdminResponse := &dto.CreateAdminResponse{
			Email: updatedUser.Email,

			NamaLengkap: updatedUser.NamaLengkap,

			Alamat:    updatedUser.Alamat,
			NoTelepon: updatedUser.NoTelepon,

			FotoProfil: updatedUser.FotoProfil,
		}

		return updateAdminResponse, nil
	}

	// Create the new student record
	CreatedUser, err := s.AdminRepo.CreateAdmin(&NewAdmin)
	if err != nil {
		return nil, err
	}

	CreateAdminResponse := &dto.CreateAdminResponse{
		Email: CreatedUser.Email,

		NamaLengkap: CreatedUser.NamaLengkap,

		Alamat:    CreatedUser.Alamat,
		NoTelepon: CreatedUser.NoTelepon,

		FotoProfil: CreatedUser.FotoProfil,
	}

	return CreateAdminResponse, nil
}
