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

	response := &dto.LoginResponse{Token: token}

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
		errDeleteImage := pkg.DeleteImage(dataDeleted[idxUser].FotoProfil)
		if errDeleteImage != nil {
			return nil, errDeleteImage
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
