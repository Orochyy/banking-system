package service

import (
	"banking-system/dto"
	"banking-system/entity"
	"banking-system/repository"
	"github.com/mashingan/smapping"
	"log"
)

type TransactionService interface {
	CreateTransaction(transactionCreateDTO dto.TransactionCreateDTO) entity.Transaction
	FindAllTransactionsByID(accountID uint64) []entity.Transaction
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepo,
	}
}

func (s *transactionService) CreateTransaction(transactionCreateDTO dto.TransactionCreateDTO) entity.Transaction {
	transaction := entity.Transaction{}
	err := smapping.FillStruct(&transaction, smapping.MapFields(&transactionCreateDTO))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.transactionRepository.CreateTransaction(transaction)
	return res
}

func (s *transactionService) FindAllTransactionsByID(accountID uint64) []entity.Transaction {
	return s.transactionRepository.FindAllTransactionsByID(accountID)
}
