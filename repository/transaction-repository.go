package repository

import (
	"banking-system/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) entity.Transaction
	FindAllTransactionsByID(accountID uint64) []entity.Transaction
}

type transactionConnection struct {
	connection *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: dbConn,
	}
}

func (db *transactionConnection) CreateTransaction(transaction entity.Transaction) entity.Transaction {
	db.connection.Save(&transaction)
	db.connection.Preload("User").Find(&transaction)
	return transaction
}

func (db *transactionConnection) FindAllTransactionsByID(accountID uint64) []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Preload("User").Where("account_sender = ?", accountID).Find(&transactions)
	return transactions
}
