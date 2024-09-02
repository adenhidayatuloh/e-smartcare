package entity

import (
	"esmartcare/pkg/errs"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"crypto/md5"
	"encoding/hex"
)

var jwtSecret = "SangatRahasia"

// SELECT * FROM `user` WHERE jenis_akun != request_jenis_akun and jenis_akun = 2
type User struct {
	Email            string `gorm:"primaryKey"`
	Password         string
	JenisAkun        string
	RequestJenisAkun string
	Password2        string
}

type ResultsJoinUsers struct {
	Email      string
	FotoProfil string
}

// TableName method sets the table name to `user`
func (User) TableName() string {
	return "users"
}

func (u *User) HashPassword() errs.MessageErr {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errs.NewInternalServerError("Failed to hash password")
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) errs.MessageErr {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		fmt.Println("Password doesn't match:", err)
		return errs.NewBadRequest("Password is not valid!")
	}

	return nil
}

func (u *User) CreateToken() (string, errs.MessageErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{

			"email": u.Email,
			"role":  u.JenisAkun,
		})

	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error:", err.Error())
		return "", errs.NewInternalServerError("Failed to sign jwt token")
	}

	return signedString, nil
}

func (u *User) ParseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthenticated("Token method is not valid")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		fmt.Print("Hallo ", err)
		return nil, errs.NewUnauthenticated("Token is not valid")
	}

	return token, nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	if isBearer := strings.HasPrefix(bearerToken, "Bearer"); !isBearer {
		return errs.NewUnauthenticated("Token type should be Bearer")
	}

	splitToken := strings.Fields(bearerToken)
	if len(splitToken) != 2 {
		return errs.NewUnauthenticated("Token is not valid")
	}

	tokenString := splitToken[1]
	token, err := u.ParseToken(tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var mapClaims jwt.MapClaims

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {

		return errs.NewUnauthenticated("Token is not valid")
	}
	mapClaims = claims

	return u.bindTokenToUserEntity(mapClaims)
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {

	email, okForEmail := claim["email"].(string)
	role, okForRole := claim["role"].(string)

	if !okForEmail {
		return errs.NewUnauthenticated("Token doesn't contains email")
	}

	if !okForRole {
		return errs.NewUnauthenticated("Token doesn't contains role")
	}
	u.Email = email
	u.JenisAkun = role

	return nil
}

// Fungsi untuk mengenkripsi password menggunakan MD5
func EncryptPasswordMD5(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// Fungsi untuk membandingkan password yang diinput dengan password yang sudah dienkripsi
func ComparePasswordMD5(inputPassword, hashedPassword string) bool {
	encryptedInputPassword := EncryptPasswordMD5(inputPassword)
	return encryptedInputPassword == hashedPassword
}
