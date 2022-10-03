package repository

import (
	"banking-system/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	InsertTransaction(t entity.Transaction) entity.Transaction
}

type transactionConnection struct {
	connection *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: dbConn,
	}
}

func (db *transactionConnection) InsertTransaction(t entity.Transaction) entity.Transaction {
	db.connection.Save(&t)
	db.connection.Preload("Account").Find(&t)
	return t
}
