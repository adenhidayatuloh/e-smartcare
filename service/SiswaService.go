package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"

	siswarepository "esmartcare/repository/siswaRepository"

	"github.com/gin-gonic/gin"
)

type SiswaService interface {
	UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreateSiswaResponse, errs.MessageErr)
	CreateOrUpdateSiswa(email string, payload *dto.CreateSiswaRequest) (*dto.CreateSiswaResponse, errs.MessageErr)
	GetAllSiswaWithPemeriksaan() ([]entity.Siswa_pemeriksaan, errs.MessageErr)
	GetSiswa(email string) (*dto.CreateSiswaResponse, errs.MessageErr)
}

type siswaService struct {
	siswaRepo siswarepository.SiswaRepository
}

// GetSiswa implements SiswaService.
func (s *siswaService) GetSiswa(email string) (*dto.CreateSiswaResponse, errs.MessageErr) {
	siswa, err := s.siswaRepo.GetSiswaByEmail(email)

	if err != nil {
		return nil, err
	}

	siswaDTO := dto.CreateSiswaResponse{
		Email:        siswa.Email,
		NIS:          siswa.NIS,
		NamaLengkap:  siswa.NamaLengkap,
		TempatLahir:  siswa.TempatLahir,
		TanggalLahir: siswa.TanggalLahir,
		Alamat:       siswa.Alamat,
		NoTelepon:    siswa.NoTelepon,
		Kelas:        siswa.Kelas,
		Agama:        siswa.Agama,
		FotoProfil:   siswa.FotoProfil,
	}

	return &siswaDTO, nil
}

// UpdateProfilPhoto implements SiswaService.
func (s *siswaService) UpdateProfilPhoto(email string, ctx *gin.Context) (*dto.CreateSiswaResponse, errs.MessageErr) {

	oldSiswa, checkEmail := s.siswaRepo.GetSiswaByEmail(email)

	if checkEmail != nil {
		return nil, errs.NewBadRequest("Please add email first")
	}

	urlImage, err := pkg.UploadImage("foto_profil", oldSiswa.Email, ctx)
	// Di sini logic nya

	if err != nil {
		return nil, err
	}

	if *urlImage == "" {
		return nil, errs.NewBadRequest("Image not detected")
	}

	// urlImageNew = strings.Replace(*urlImage, "-temp", "", -1)

	// if oldAdmin.FotoProfil != "" {
	// 	// Delete the old image only after the new image is uploaded successfully
	// 	errDeleteImage := pkg.DeleteImage(oldAdmin.FotoProfil)
	// 	if errDeleteImage != nil {
	// 		return nil, errDeleteImage
	// 	}
	// }

	// // // Rename the new image from temporary to final name
	// err = pkg.RenameImage(*urlImage, urlImageNew)
	// if err != nil {
	// 	return nil, errs.NewInternalServerError("Error on upload image")
	// }

	Newsiswa := entity.Siswa{
		Email: email,

		FotoProfil: *urlImage,
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
