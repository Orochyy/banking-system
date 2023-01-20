package controller

import (
	"banking-system/dto"
	"banking-system/entity"
	"banking-system/helper"
	"banking-system/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransactionController interface {
	CreateTransaction(context *gin.Context)
	GetAllTransactionByAccountID(context *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
	accountService     service.AccountService
	jwtService         service.JWTService
}

func NewTransactionController(transactionService service.TransactionService, accountService service.AccountService, jwtServ service.JWTService) TransactionController {
	return &transactionController{
		transactionService: transactionService,
		accountService:     accountService,
		jwtService:         jwtServ,
	}
}

func (c *transactionController) CreateTransaction(context *gin.Context) {
	var transactionCreateDTO dto.TransactionCreateDTO
	errDTO := context.ShouldBind(&transactionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		transactionCreateDTO.UserID = convertedUserID
	}

	if transactionCreateDTO.Type != "withdraw" && transactionCreateDTO.Type != "deposit" && transactionCreateDTO.Type != "transfer" {
		res := helper.BuildErrorResponse("Failed to process request", "Invalid transaction type", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	currencyAccountFrom := c.accountService.GetCurrency(transactionCreateDTO.AccountSender)
	currencyAccountTo := c.accountService.GetCurrency(transactionCreateDTO.AccountRecipient)

	if currencyAccountFrom != currencyAccountTo {
		res := helper.BuildErrorResponse("Currency not supported", "Currency not supported", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := c.transactionService.CreateTransaction(transactionCreateDTO)

	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	hexSender := c.accountService.GetHex(transactionCreateDTO.AccountSender)
	hexRecipient := c.accountService.GetHex(transactionCreateDTO.AccountRecipient)

	if hexSender == "" || hexRecipient == "" {
		res := helper.BuildErrorResponse("Failed to process request", "Invalid account", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	accountSender := c.accountService.FindByHex(hexSender)
	accountRecipient := c.accountService.FindByHex(hexRecipient)

	if accountSender.Amount < transactionCreateDTO.Amount {
		res := helper.BuildErrorResponse("Failed to process request", "Insufficient funds", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	accountSender.Amount = accountSender.Amount - transactionCreateDTO.Amount
	accountRecipient.Amount = accountRecipient.Amount + transactionCreateDTO.Amount

	c.accountService.UpdateAmountAccount(accountSender)
	c.accountService.UpdateAmountAccount(accountRecipient)

	response := helper.BuildResponse(true, "OK", res)
	context.JSON(http.StatusCreated, response)
}

func (c *transactionController) GetAllTransactionByAccountID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	for _, transaction := range c.transactionService.FindAllTransactionsByID(id) {
		if (transaction == entity.Transaction{}) {
			res := helper.BuildErrorResponse("No transaction found with given id", "No data found", helper.EmptyObj{})
			context.JSON(http.StatusNotFound, res)
		} else {
			res := helper.BuildResponse(true, "OK", transaction)
			context.JSON(http.StatusOK, res)
		}
	}
}

func (c *transactionController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
