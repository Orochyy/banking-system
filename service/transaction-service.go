package service

import (
	"banking-system/dto"
	"banking-system/entity"
	"banking-system/repository"
	"github.com/mashingan/smapping"
	"log"
)

type TransactionService interface {
	Insert(t dto.TransactionCreateDTO) entity.Transaction
	All() []entity.Transaction
	FindByID(transactionID uint64) entity.Transaction
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepo,
	}
}

func (service *transactionService) Insert(t dto.TransactionCreateDTO) entity.Transaction {
	transaction := entity.Transaction{}
	err := smapping.FillStruct(&transaction, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.transactionRepository.InsertTransaction(transaction)
	return res
}

func (service *transactionService) All() []entity.Transaction {
	return service.transactionRepository.AllTransactions()
}

func (service *transactionService) FindByID(transactionID uint64) entity.Transaction {
	return service.transactionRepository.FindTransactionByID(transactionID)
}
