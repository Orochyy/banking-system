package service

import (
	"banking-system/dto"
	"banking-system/entity"
	"banking-system/repository"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type AccountService interface {
	Insert(b dto.AccountCreateDTO) entity.Account
	Update(b dto.AccountUpdateDTO) entity.Account
	Delete(b entity.Account)
	All() []entity.Account
	FindByID(accountID uint64) entity.Account
	IsAllowedToEdit(userID string, accountID uint64) bool
	GetCurrency(accountID uint64) string
	CreateHex() string
	UpdateAmountAccount(a entity.Account) entity.Account
	GetHex(accountID uint64) string
	FindByHex(hex string) entity.Account
	ParseString(s string) (string, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{
		accountRepository: accountRepo,
	}
}

func (service *accountService) Insert(b dto.AccountCreateDTO) entity.Account {
	account := entity.Account{}
	err := smapping.FillStruct(&account, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.accountRepository.InsertAccount(account)
	return res
}

func (service *accountService) Update(b dto.AccountUpdateDTO) entity.Account {
	account := entity.Account{}
	err := smapping.FillStruct(&account, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.accountRepository.UpdateAccount(account)
	return res
}

func (service *accountService) Delete(b entity.Account) {
	service.accountRepository.DeleteAccount(b)
}

func (service *accountService) All() []entity.Account {
	return service.accountRepository.AllAccounts()
}

func (service *accountService) FindByID(accountID uint64) entity.Account {
	return service.accountRepository.FindAccountByID(accountID)
}

func (service *accountService) IsAllowedToEdit(userID string, accountID uint64) bool {
	b := service.accountRepository.FindAccountByID(accountID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}

func (service *accountService) GetCurrency(accountID uint64) string {
	b := service.accountRepository.FindAccountByID(accountID)
	return b.Currency
}

func (service *accountService) UpdateAmountAccount(a entity.Account) entity.Account {
	account := entity.Account{}
	err := smapping.FillStruct(&account, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.accountRepository.UpdateAccount(account)

	return res
}

func (service *accountService) CreateHex() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (service *accountService) GetHex(accountID uint64) string {
	b := service.accountRepository.FindAccountByID(accountID)
	return b.Hex
}

func (service *accountService) FindByHex(hex string) entity.Account {
	return service.accountRepository.FindAccountByHex(hex)
}

func (service *accountService) ParseString(s string) (string, error) {
	var result string
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return "", err
	}
	return result, nil
}
