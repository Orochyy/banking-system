package repository

import (
	"banking-system/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	InsertTransaction(t entity.Transaction) entity.Transaction
	UpdateTransaction(t entity.Transaction) entity.Transaction
	GetTransaction(t entity.Transaction) entity.Transaction
	FindTransactionByID(transactionID uint64) entity.Transaction
	AllTransactions() []entity.Transaction
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

func (db *transactionConnection) UpdateTransaction(t entity.Transaction) entity.Transaction {
	db.connection.Save(&t)
	db.connection.Preload("Account").Find(&t)
	return t
}

func (db *transactionConnection) GetTransaction(t entity.Transaction) entity.Transaction {
	db.connection.Preload("Account").Find(&t)
	return t
}

func (db *transactionConnection) FindTransactionByID(transactionID uint64) entity.Transaction {
	var transaction entity.Transaction
	db.connection.Preload("Account").Find(&transaction, transactionID)
	return transaction
}

func (db *transactionConnection) AllTransactions() []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Preload("Account").Find(&transactions)
	return transactions
}
