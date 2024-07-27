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
		Select("user.email, " + joinTable + ".foto_profil").
		Joins("left join " + joinTable + " on " + joinTable + ".email = user.email").
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

	// if err := db.Where("JenisAkun != RequestJenisAkun").Find(&users).Error; err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//     return
	// }

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
