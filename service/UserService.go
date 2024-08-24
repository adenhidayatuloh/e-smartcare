package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"
	"fmt"
	"strconv"

	"esmartcare/repository/userrepository"
)

type UserService interface {
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	GetAllUsers(jenis_akun string) ([]dto.GetAllUsersResponse, errs.MessageErr)
	GetAllUsersNotValidate(jenis_akun string) ([]dto.GetAllUsersResponse, errs.MessageErr)
	UpdateUser(email string) (*dto.UpdateUserResponse, errs.MessageErr)
	DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr)
	GetAllDataUser(jenisAkun string) (interface{}, errs.MessageErr)
}

type userService struct {
	userRepo userrepository.UserRepository
}

func NewUserService(userRepo userrepository.UserRepository) UserService {
	return &userService{userRepo}
}

func (u *userService) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user := entity.User{}

	if payload.JenisAkun == "siswa" {
		user = entity.User{

			Email:            payload.Email,
			Password:         payload.Password,
			JenisAkun:        "3",
			RequestJenisAkun: "3",
		}
	} else if payload.JenisAkun == "admin" {
		user = entity.User{

			Email:            payload.Email,
			Password:         payload.Password,
			RequestJenisAkun: "1",
		}

	} else if payload.JenisAkun == "pakar" {
		user = entity.User{

			Email:            payload.Email,
			Password:         payload.Password,
			RequestJenisAkun: "2",
		}

	}

	_, checkEmail := u.userRepo.GetUserByEmail(user.Email)

	if checkEmail == nil {
		return nil, errs.NewBadRequest("email already exists")
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	registeredUser, err := u.userRepo.Register(&user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		Email:     registeredUser.Email,
		JenisAkun: payload.JenisAkun,
	}

	return response, nil
}

func (u *userService) Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if user.JenisAkun == "" {
		return nil, errs.NewBadRequest("Account has not been validated")
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		return nil, err
	}

	token, err2 := user.CreateToken()
	if err2 != nil {
		return nil, err2
	}

	role := ""

	if user.JenisAkun == "1" {
		role = "admin"

	} else if user.JenisAkun == "2" {
		role = "pakar"
	} else {
		role = "siswa"
	}

	response := &dto.LoginResponse{Token: token, Role: role}

	return response, nil
}

// GetAllUser implements UserService.
func (u *userService) GetAllUsers(jenis_akun string) ([]dto.GetAllUsersResponse, errs.MessageErr) {

	if jenis_akun != "" {
		jenis_akun_int, err := strconv.Atoi(jenis_akun)

		if err != nil {
			return nil, errs.NewBadRequest("jenis_akun must int")
		}

		if !(jenis_akun_int >= 1 && jenis_akun_int <= 3) {
			return nil, errs.NewBadRequest("jenis_akun must be 1 (admin), 2 (pakar), or 3 (siswa)")
		}

	}

	AllUsers, errGetUser := u.userRepo.GetAllUsers(jenis_akun)

	if errGetUser != nil {
		return nil, errGetUser
	}

	AllUsersDto := []dto.GetAllUsersResponse{}

	for _, eachUser := range AllUsers {

		User := dto.GetAllUsersResponse{
			Email:            eachUser.Email,
			JenisAkun:        eachUser.JenisAkun,
			RequestJenisAkun: eachUser.RequestJenisAkun,
		}

		AllUsersDto = append(AllUsersDto, User)
	}

	return AllUsersDto, nil

}

// GetAllUsersNotValidate implements UserService.
func (u *userService) GetAllUsersNotValidate(jenis_akun string) ([]dto.GetAllUsersResponse, errs.MessageErr) {

	if jenis_akun != "" {
		jenis_akun_int, err := strconv.Atoi(jenis_akun)

		if err != nil {
			return nil, errs.NewBadRequest("jenis_akun must int")
		}

		if !(jenis_akun_int >= 1 && jenis_akun_int <= 2) {
			return nil, errs.NewBadRequest("jenis_akun must be 1 (admin), 2 (pakar)")
		}

	}
	AllUsers, errGetUser := u.userRepo.GetAllUsersNotValidate(jenis_akun)

	if errGetUser != nil {
		return nil, errGetUser
	}

	AllUsersDto := []dto.GetAllUsersResponse{}

	for _, eachUser := range AllUsers {

		User := dto.GetAllUsersResponse{
			Email:            eachUser.Email,
			JenisAkun:        eachUser.JenisAkun,
			RequestJenisAkun: eachUser.RequestJenisAkun,
		}

		AllUsersDto = append(AllUsersDto, User)
	}

	return AllUsersDto, nil
}

func (u *userService) UpdateUser(email string) (*dto.UpdateUserResponse, errs.MessageErr) {

	//err := pkg.ValidateStruct(payload)

	oldUser, err := u.userRepo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if oldUser.JenisAkun != "" {
		return nil, errs.NewBadRequest("Email has validated")
	}

	newUser := entity.User{}

	if oldUser.RequestJenisAkun == "1" {
		newUser.JenisAkun = "1"
	} else if oldUser.RequestJenisAkun == "2" {
		newUser.JenisAkun = "2"
	}

	updatedUser, err := u.userRepo.UpdateUser(oldUser, &newUser)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		Message: fmt.Sprintf("Sucess validate user email = %s", updatedUser.Email),
	}

	return response, nil
}

func (u *userService) DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr) {

	tableJoin := ""

	OldUser, err := u.userRepo.GetUserByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if OldUser.JenisAkun == "1" {
		tableJoin = "admin"
	} else if OldUser.JenisAkun == "2" {
		tableJoin = "pakar"
	} else if OldUser.JenisAkun == "3" {
		tableJoin = "siswa"
	}

	dataDeleted, err := u.userRepo.GetUserJoin(tableJoin)

	if err != nil {
		return nil, err
	}

	idxUser := 0

	for i, v := range dataDeleted {
		if OldUser.Email == v.Email {
			idxUser = i
		}
	}

	if dataDeleted[idxUser].FotoProfil != "" {
		// Delete the old image only after the new image is uploaded successfully
		errDeleteImage := pkg.DeleteImage(dataDeleted[idxUser].Email)
		if errDeleteImage != nil {
			return nil, err
		}
	}

	if err = u.userRepo.DeleteUser(user); err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "Your account has been successfully deleted",
	}

	return response, nil
}

func (s *userService) GetAllDataUser(jenisAkun string) (interface{}, errs.MessageErr) {
	data, err := s.userRepo.GetAllDataUser(jenisAkun)
	if err != nil {
		return nil, err
	}

	allAdminPakarSiswa := make([]interface{}, 0)

	switch jenisAkun {
	case "1": // Admin and Pakar
		adminPakarResponse := make([]dto.GetAdminPakarResponse, 0)

		fmt.Print(data.([]entity.Admin))
		for _, adminPakar := range data.([]entity.Admin) {
			response := dto.GetAdminPakarResponse{
				Email:            adminPakar.Email,
				NamaLengkap:      adminPakar.NamaLengkap,
				Alamat:           adminPakar.Alamat,
				NoTelepon:        adminPakar.NoTelepon,
				FotoProfil:       adminPakar.FotoProfil,
				JenisAkun:        adminPakar.User.JenisAkun,
				RequestJenisAkun: adminPakar.User.RequestJenisAkun,
			}
			adminPakarResponse = append(adminPakarResponse, response)
		}
		return adminPakarResponse, nil

	case "2": // Pakar
		adminPakarResponse := make([]dto.GetAdminPakarResponse, 0)
		for _, pakar := range data.([]entity.Pakar) {
			response := dto.GetAdminPakarResponse{
				Email:            pakar.Email,
				NamaLengkap:      pakar.NamaLengkap,
				Alamat:           pakar.Alamat,
				NoTelepon:        pakar.NoTelepon,
				FotoProfil:       pakar.FotoProfil,
				JenisAkun:        pakar.User.JenisAkun,
				RequestJenisAkun: pakar.User.RequestJenisAkun,
			}
			adminPakarResponse = append(adminPakarResponse, response)
		}
		return adminPakarResponse, nil

	case "3": // Siswa
		siswaResponse := make([]dto.GetSiswaResponse, 0)
		for _, siswa := range data.([]entity.Siswa) {
			response := dto.GetSiswaResponse{
				Email:            siswa.Email,
				NIS:              siswa.NIS,
				NamaLengkap:      siswa.NamaLengkap,
				TempatLahir:      siswa.TempatLahir,
				TanggalLahir:     siswa.TanggalLahir,
				Alamat:           siswa.Alamat,
				NoTelepon:        siswa.NoTelepon,
				Kelas:            siswa.Kelas,
				Agama:            siswa.Agama,
				FotoProfil:       siswa.FotoProfil,
				JenisAkun:        siswa.User.JenisAkun,
				RequestJenisAkun: siswa.User.RequestJenisAkun,
			}
			siswaResponse = append(siswaResponse, response)
		}
		return siswaResponse, nil

	}

	allAdmin := make([]dto.GetAdminPakarResponse, 0)

	admin := data.(map[string]interface{})["admin"].([]entity.Admin)
	for _, adminPakar := range admin {

		response := dto.GetAdminPakarResponse{
			Email:            adminPakar.Email,
			NamaLengkap:      adminPakar.NamaLengkap,
			Alamat:           adminPakar.Alamat,
			NoTelepon:        adminPakar.NoTelepon,
			FotoProfil:       adminPakar.FotoProfil,
			JenisAkun:        adminPakar.User.JenisAkun,
			RequestJenisAkun: adminPakar.User.RequestJenisAkun,
		}
		allAdmin = append(allAdmin, response)
	}

	allPakar := make([]dto.GetAdminPakarResponse, 0)

	pakar := data.(map[string]interface{})["pakar"].([]entity.Pakar)
	for _, pakar := range pakar {
		response := dto.GetAdminPakarResponse{
			Email:            pakar.Email,
			NamaLengkap:      pakar.NamaLengkap,
			Alamat:           pakar.Alamat,
			NoTelepon:        pakar.NoTelepon,
			FotoProfil:       pakar.FotoProfil,
			JenisAkun:        pakar.User.JenisAkun,
			RequestJenisAkun: pakar.User.RequestJenisAkun,
		}
		allPakar = append(allPakar, response)
	}

	allSiswa := make([]dto.GetSiswaResponse, 0)

	siswa := data.(map[string]interface{})["siswa"].([]entity.Siswa)
	for _, siswa := range siswa {
		response := dto.GetSiswaResponse{
			Email:            siswa.Email,
			NIS:              siswa.NIS,
			NamaLengkap:      siswa.NamaLengkap,
			TempatLahir:      siswa.TempatLahir,
			TanggalLahir:     siswa.TanggalLahir,
			Alamat:           siswa.Alamat,
			NoTelepon:        siswa.NoTelepon,
			Kelas:            siswa.Kelas,
			Agama:            siswa.Agama,
			FotoProfil:       siswa.FotoProfil,
			JenisAkun:        siswa.User.JenisAkun,
			RequestJenisAkun: siswa.User.RequestJenisAkun,
		}
		allSiswa = append(allSiswa, response)
	}

	allAdminPakarSiswa = append(allAdminPakarSiswa, allAdmin, allPakar, allSiswa)

	return allAdminPakarSiswa, nil
}
