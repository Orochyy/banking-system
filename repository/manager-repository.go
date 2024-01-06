package repository

import (
	"banking-system/entity"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

type ManagerRepository interface {
	InsertPassword(p entity.Password) entity.Password
	UpdatePassword(p entity.Password) entity.Password
	DeletePassword(p entity.Password)
	GetAll() []entity.Password
}

type managerConnection struct {
	connection *gorm.DB
}

func NewManagerRepository(dbConn *gorm.DB) ManagerRepository {
	return &managerConnection{
		connection: dbConn,
	}
}

func (db *managerConnection) InsertPassword(b entity.Password) entity.Password {
	secondCrypt := hashAndSalt([]byte(b.Password))
	b.Password = doubleHashAndSaltWithHMAC([]byte(secondCrypt))
	db.connection.Save(&b)
	return b
}

func (db *managerConnection) UpdatePassword(b entity.Password) entity.Password {
	db.connection.Save(&b)
	db.connection.Preload("Password").Find(&b)
	return b
}

func (db *managerConnection) DeletePassword(b entity.Password) {
	db.connection.Delete(&b)
}

func (db *managerConnection) GetAll() []entity.Password {
	var manager []entity.Password
	db.connection.Preload("Password").Find(&manager)
	return manager
}

func hashWithHMAC(input string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(input))
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	fmt.Print("hash2:", hash)
	return hash
}

func doubleHashAndSaltWithHMAC(pwd []byte) string {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	secretKey := os.Getenv("SECRET_KEY")
	firstRoundHash := hashAndSalt(pwd)

	secondRoundHash := hashWithHMAC(firstRoundHash, secretKey)
	fmt.Println("secret for crypt", secretKey)
	fmt.Println("hash3:", secondRoundHash)
	return secondRoundHash
}
