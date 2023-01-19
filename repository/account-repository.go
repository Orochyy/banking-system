package repository

import (
	"banking-system/entity"
	"gorm.io/gorm"
)

type AccountRepository interface {
	InsertAccount(b entity.Account) entity.Account
	UpdateAccount(b entity.Account) entity.Account
	DeleteAccount(b entity.Account)
	AllAccounts() []entity.Account
	FindAccountByID(accountID uint64) entity.Account
	GetHex(accountID uint64) string
	FindAccountByHex(hex string) entity.Account
}

type accountConnection struct {
	connection *gorm.DB
}

func NewAccountRepository(dbConn *gorm.DB) AccountRepository {
	return &accountConnection{
		connection: dbConn,
	}
}

func (db *accountConnection) InsertAccount(b entity.Account) entity.Account {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *accountConnection) UpdateAccount(b entity.Account) entity.Account {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *accountConnection) DeleteAccount(b entity.Account) {
	db.connection.Delete(&b)
}

func (db *accountConnection) FindAccountByID(accountID uint64) entity.Account {
	var account entity.Account
	db.connection.Preload("User").Find(&account, accountID)
	return account
}

func (db *accountConnection) AllAccounts() []entity.Account {
	var accounts []entity.Account
	db.connection.Preload("User").Find(&accounts)
	return accounts
}

func (db *accountConnection) GetHex(accountID uint64) string {
	var account entity.Account
	db.connection.Preload("User").Find(&account, accountID)
	return account.Hex
}

func (db *accountConnection) FindAccountByHex(hex string) entity.Account {
	var account entity.Account
	db.connection.Where("hex = ?", hex).Find(&account)
	return account
}
