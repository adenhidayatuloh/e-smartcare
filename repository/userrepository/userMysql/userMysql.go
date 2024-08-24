package usermysql

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	"esmartcare/repository/userrepository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userMySql struct {
	db *gorm.DB
}

// GetUserJoin implements userrepository.UserRepository.
func (u *userMySql) GetUserJoin(joinTable string) ([]entity.ResultsJoinUsers, errs.MessageErr) {

	var results []entity.ResultsJoinUsers

	err := u.db.Model(&entity.User{}).
		Select("users.email, " + joinTable + ".foto_profil").
		Joins("left join " + joinTable + " on " + joinTable + ".email = users.email").
		Scan(&results).Error

	if err != nil {
		return nil, errs.NewInternalServerError("Error fetching user join data")
	}

	return results, nil
}

func NewUserMySql(db *gorm.DB) userrepository.UserRepository {
	return &userMySql{db}
}

func (u *userMySql) Register(user *entity.User) (*entity.User, errs.MessageErr) {

	if err := u.db.Create(user).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to register new user")
	}

	return user, nil
}

func (u *userMySql) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}

func (u *userMySql) GetUserByID(id uint) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with id %d is not found", id))
	}

	return &user, nil
}

// GetAllUser implements userrepository.UserRepository.
func (u *userMySql) GetAllUsers(jenis_akun string) ([]entity.User, errs.MessageErr) {
	var users []entity.User
	query := u.db.Model(&entity.User{})
	if jenis_akun != "" {
		query = query.Where("jenis_akun = ?", jenis_akun)
		if err := query.Find(&users).Error; err != nil {
			return nil, errs.NewNotFound(fmt.Sprintf("Users with jenis akun %s is not found", jenis_akun))

		}
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, errs.NewNotFound("Users with not found")

	}

	return users, nil

}

// GetAllUserNotValidate implements userrepository.UserRepository.
func (u *userMySql) GetAllUsersNotValidate(jenis_akun string) ([]entity.User, errs.MessageErr) {
	var users []entity.User
	query := u.db.Model(&entity.User{})
	if jenis_akun != "" {
		query = query.Where("request_jenis_akun = ? AND jenis_akun != request_jenis_akun", jenis_akun)
	} else {
		query = query.Where("jenis_akun != request_jenis_akun")
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, errs.NewNotFound("Users which not validate not found")

	}

	return users, nil
}

func (u *userMySql) UpdateUser(oldUser *entity.User, newUser *entity.User) (*entity.User, errs.MessageErr) {
	if err := u.db.Model(oldUser).Updates(newUser).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update user with email %s", oldUser.Email))
	}

	return oldUser, nil
}

func (u *userMySql) DeleteUser(user *entity.User) errs.MessageErr {
	if err := u.db.Delete(user).Error; err != nil {
		log.Println(err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete user email %s", user.Email))
	}

	return nil
}

func (u *userMySql) GetAllDataUser(jenis_akun string) (interface{}, errs.MessageErr) {
	var result interface{}

	switch jenis_akun {
	case "1":
		var allAdmin []entity.Admin
		if err := u.db.Preload("User").Order("email ASC").Find(&allAdmin).Error; err != nil {
			return nil, errs.NewNotFound("Admins not found")
		}

		fmt.Println(allAdmin)
		result = allAdmin

	case "2":
		var allPakar []entity.Pakar
		if err := u.db.Preload("User").Order("email ASC").Find(&allPakar).Error; err != nil {
			return nil, errs.NewNotFound("pakar not found")
		}
		result = allPakar

	case "3":
		var allSiswa []entity.Siswa
		if err := u.db.Preload("User").Order("email ASC").Find(&allSiswa).Error; err != nil {
			return nil, errs.NewNotFound("Students not found")
		}
		result = allSiswa

	default:
		// Mengambil semua data Admin, Pakar, dan Siswa jika jenis_akun tidak diisi
		var allAdmin []entity.Admin
		if err := u.db.Preload("User").Order("email ASC").Find(&allAdmin).Error; err != nil {
			return nil, errs.NewNotFound("Admins not found")
		}

		var allPakar []entity.Pakar
		if err := u.db.Preload("User").Order("email ASC").Find(&allPakar).Error; err != nil {
			return nil, errs.NewNotFound("Experts not found")
		}

		var allSiswa []entity.Siswa
		if err := u.db.Preload("User").Order("email ASC").Find(&allSiswa).Error; err != nil {
			return nil, errs.NewNotFound("Students not found")
		}

		result = map[string]interface{}{
			"admin": allAdmin,
			"pakar": allPakar,
			"siswa": allSiswa,
		}
	}

	return result, nil
}
