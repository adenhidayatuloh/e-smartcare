package mysql

import (
	"errors"
	"esmartcare/entity"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
	db       *gorm.DB
	err      error
)

func GetDBConfig() gorm.Dialector {

	dbConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)
	return mysql.Open(dbConfig)
}

func GetDBInstance() *gorm.DB {
	return db
}

func seedAdmin() {
	admin := &entity.User{
		Email:            "adminku@gmail.com",
		Password:         "admin12345",
		JenisAkun:        "1",
		RequestJenisAkun: "1",
	}
	err := admin.HashPassword()
	if err != nil {
		fmt.Println(err)
	}

	err2 := db.Create(admin).Error
	if err2 != nil {
		fmt.Println(err2)
	}

	log.Println("Admin account seed success!")
}

func init() {
	db, err = gorm.Open(GetDBConfig(), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := db.First(&entity.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		seedAdmin()
	}

	log.Println("Connected to DB!")
}
