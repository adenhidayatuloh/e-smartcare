package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"
	"strings"

	siswarepository "esmartcare/repository/siswaRepository"

	"github.com/gin-gonic/gin"
)

type SiswaService interface {
	CreateOrUpdateSiswaWithProfilPhoto(email string, payload *dto.CreateSiswaRequest, ctx *gin.Context) (*dto.CreateSiswaResponse, errs.MessageErr)
	UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreateSiswaResponse, errs.MessageErr)
	CreateOrUpdateSiswa(email string, payload *dto.CreateSiswaRequest) (*dto.CreateSiswaResponse, errs.MessageErr)
	GetAllSiswaWithPemeriksaan() ([]entity.Siswa_pemeriksaan, errs.MessageErr)

	//Update(email string, payload *dto.CreateSiswaRequestAndResponse) (*dto.CreateSiswaRequestAndResponse, errs.MessageErr)
	//UpdatePhoto(email string) (*dto.UpdatePhotoResponse, errs.MessageErr)
	// Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	// GetAllUsers(jenis_akun string) ([]dto.GetAllUsersResponse, errs.MessageErr)
	// GetAllUsersNotValidate(jenis_akun string) ([]dto.GetAllUsersResponse, errs.MessageErr)
	// UpdateUser(email string) (*dto.UpdateUserResponse, errs.MessageErr)
	//DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr)
}

type siswaService struct {
	siswaRepo siswarepository.SiswaRepository
}

// UpdateProfilPhoto implements SiswaService.
func (s *siswaService) UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreateSiswaResponse, errs.MessageErr) {

	urlImageNew := ""
	oldSiswa, checkEmail := s.siswaRepo.GetSiswaByEmail(email)

	if checkEmail != nil {
		return nil, errs.NewBadRequest("Please add email first")
	}

	urlImage, err := pkg.UploadImage("foto_profil", oldSiswa.Email, ctx)

	if err != nil {
		return nil, err
	}

	if *urlImage == "" {
		return nil, errs.NewBadRequest("Image not detected")
	}

	urlImageNew = strings.Replace(*urlImage, "-temp", "", -1)

	if oldSiswa.FotoProfil != "" {
		// Delete the old image only after the new image is uploaded successfully
		errDeleteImage := pkg.DeleteImage(oldSiswa.FotoProfil)
		if errDeleteImage != nil {
			return nil, errDeleteImage
		}
	}

	// Rename the new image from temporary to final name
	err = pkg.RenameImage(*urlImage, urlImageNew)
	if err != nil {
		return nil, errs.NewInternalServerError("Error on upload image")
	}

	Newsiswa := entity.Siswa{
		Email: email,

		FotoProfil: urlImageNew,
	}

	// Update the student record
	updatedUser, err := s.siswaRepo.UpdateSiswa(oldSiswa, &Newsiswa)
	if err != nil {
		return nil, errs.NewBadRequest("Cannot update siswa")
	}

	updateSiswaResponse := &dto.CreateSiswaResponse{
		Email:        updatedUser.Email,
		NIS:          updatedUser.NIS,
		NamaLengkap:  updatedUser.NamaLengkap,
		TempatLahir:  updatedUser.TempatLahir,
		TanggalLahir: updatedUser.TanggalLahir,
		Alamat:       updatedUser.Alamat,
		NoTelepon:    updatedUser.NoTelepon,
		Kelas:        updatedUser.Kelas,
		Agama:        updatedUser.Agama,
		FotoProfil:   updatedUser.FotoProfil,
	}

	return updateSiswaResponse, nil

}

func NewSiswaService(siswaRepo siswarepository.SiswaRepository) SiswaService {
	return &siswaService{siswaRepo}
}

func (s *siswaService) CreateOrUpdateSiswaWithProfilPhoto(email string, payload *dto.CreateSiswaRequest, ctx *gin.Context) (*dto.CreateSiswaResponse, errs.MessageErr) {

	urlImageNew := ""
	Newsiswa := entity.Siswa{
		Email:        email,
		NIS:          payload.NIS,
		NamaLengkap:  payload.NamaLengkap,
		TempatLahir:  payload.TempatLahir,
		TanggalLahir: payload.TanggalLahir,
		Alamat:       payload.Alamat,
		NoTelepon:    payload.NoTelepon,
		Kelas:        payload.Kelas,
		Agama:        payload.Agama,
		FotoProfil:   payload.FotoProfil,
	}

	oldSiswa, checkEmail := s.siswaRepo.GetSiswaByEmail(email)

	// Upload image with a temporary name
	urlImage, err := pkg.UploadImage("foto_profil", oldSiswa.Email, ctx)

	if err != nil {
		return nil, errs.NewBadRequest("Cannot upload foto profil siswa")
	}

	if *urlImage != "" {
		urlImageNew = strings.Replace(*urlImage, "-temp", "", -1)
	}

	if checkEmail == nil {
		if oldSiswa.FotoProfil != "" && *urlImage != "" {
			// Delete the old image only after the new image is uploaded successfully
			errDeleteImage := pkg.DeleteImage(oldSiswa.FotoProfil)
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
			Newsiswa.FotoProfil = urlImageNew
		}

		// Update the student record
		updatedUser, err := s.siswaRepo.UpdateSiswa(oldSiswa, &Newsiswa)
		if err != nil {
			return nil, errs.NewBadRequest("Cannot update siswa")
		}
		updateSiswaResponse := &dto.CreateSiswaResponse{
			Email:        updatedUser.Email,
			NIS:          updatedUser.NIS,
			NamaLengkap:  updatedUser.NamaLengkap,
			TempatLahir:  updatedUser.TempatLahir,
			TanggalLahir: updatedUser.TanggalLahir,
			Alamat:       updatedUser.Alamat,
			NoTelepon:    updatedUser.NoTelepon,
			Kelas:        updatedUser.Kelas,
			Agama:        updatedUser.Agama,
			FotoProfil:   updatedUser.FotoProfil,
		}

		return updateSiswaResponse, nil
	}

	if *urlImage != "" {
		// Rename the new image from temporary to final name
		err := pkg.RenameImage(*urlImage, urlImageNew)
		if err != nil {
			return nil, errs.NewInternalServerError("Cannot rename image")
		}
		Newsiswa.FotoProfil = urlImageNew
	}

	// Create the new student record
	CreatedUser, err := s.siswaRepo.CreateSiswa(&Newsiswa)
	if err != nil {
		return nil, err
	}

	CreateSiswaResponse := &dto.CreateSiswaResponse{
		Email:        CreatedUser.Email,
		NIS:          CreatedUser.NIS,
		NamaLengkap:  CreatedUser.NamaLengkap,
		TempatLahir:  CreatedUser.TempatLahir,
		TanggalLahir: CreatedUser.TanggalLahir,
		Alamat:       CreatedUser.Alamat,
		NoTelepon:    CreatedUser.NoTelepon,
		Kelas:        CreatedUser.Kelas,
		Agama:        CreatedUser.Agama,
		FotoProfil:   CreatedUser.FotoProfil,
	}

	return CreateSiswaResponse, nil
}

// CreateOrUpdateSiswa implements SiswaService.
func (s *siswaService) CreateOrUpdateSiswa(email string, payload *dto.CreateSiswaRequest) (*dto.CreateSiswaResponse, errs.MessageErr) {

	Newsiswa := entity.Siswa{
		Email:        email,
		NIS:          payload.NIS,
		NamaLengkap:  payload.NamaLengkap,
		TempatLahir:  payload.TempatLahir,
		TanggalLahir: payload.TanggalLahir,
		Alamat:       payload.Alamat,
		NoTelepon:    payload.NoTelepon,
		Kelas:        payload.Kelas,
		Agama:        payload.Agama,
	}

	oldSiswa, checkEmail := s.siswaRepo.GetSiswaByEmail(email)

	if checkEmail == nil {

		updatedUser, err := s.siswaRepo.UpdateSiswa(oldSiswa, &Newsiswa)
		if err != nil {
			return nil, err
		}
		updateSiswaResponse := &dto.CreateSiswaResponse{
			Email:        updatedUser.Email,
			NIS:          updatedUser.NIS,
			NamaLengkap:  updatedUser.NamaLengkap,
			TempatLahir:  updatedUser.TempatLahir,
			TanggalLahir: updatedUser.TanggalLahir,
			Alamat:       updatedUser.Alamat,
			NoTelepon:    updatedUser.NoTelepon,
			Kelas:        updatedUser.Kelas,
			Agama:        updatedUser.Agama,
			FotoProfil:   updatedUser.FotoProfil,
		}

		return updateSiswaResponse, nil
	}

	// Create the new student record
	CreatedUser, err := s.siswaRepo.CreateSiswa(&Newsiswa)
	if err != nil {
		return nil, err
	}

	CreateSiswaResponse := &dto.CreateSiswaResponse{
		Email:        CreatedUser.Email,
		NIS:          CreatedUser.NIS,
		NamaLengkap:  CreatedUser.NamaLengkap,
		TempatLahir:  CreatedUser.TempatLahir,
		TanggalLahir: CreatedUser.TanggalLahir,
		Alamat:       CreatedUser.Alamat,
		NoTelepon:    CreatedUser.NoTelepon,
		Kelas:        CreatedUser.Kelas,
		Agama:        CreatedUser.Agama,
		FotoProfil:   CreatedUser.FotoProfil,
	}

	return CreateSiswaResponse, nil
}

func (s *siswaService) GetAllSiswaWithPemeriksaan() ([]entity.Siswa_pemeriksaan, errs.MessageErr) {
	return s.siswaRepo.GetAllSiswaWithPemeriksaan()
}
