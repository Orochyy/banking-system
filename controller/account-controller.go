package controller

import (
	"banking-system/dto"
	"banking-system/entity"
	"banking-system/helper"
	"banking-system/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AccountController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type accountController struct {
	accountService service.AccountService
	jwtService     service.JWTService
}

func NewAccountController(accountService service.AccountService, jwtServ service.JWTService) AccountController {
	return &accountController{
		accountService: accountService,
		jwtService:     jwtServ,
	}
}

func (c *accountController) All(context *gin.Context) {
	var accounts []entity.Account = c.accountService.All()
	res := helper.BuildResponse(true, "OK", accounts)
	context.JSON(http.StatusOK, res)
}

func (c *accountController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var account entity.Account = c.accountService.FindByID(id)
	if (account == entity.Account{}) {
		res := helper.BuildErrorResponse("No account found with given id", "No data found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", account)
		context.JSON(http.StatusOK, res)
	}
}

func (c *accountController) Insert(context *gin.Context) {
	var accountCreateDTO dto.AccountCreateDTO
	errDTO := context.ShouldBind(&accountCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		if accountCreateDTO.Currency != "USD" && accountCreateDTO.Currency != "COP" && accountCreateDTO.Currency != "MXN" {
			res := helper.BuildErrorResponse("Failed to process request", "Currency not allowed", helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)
			return
		}
	}
	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		accountCreateDTO.UserID = convertedUserID
	}
	result := c.accountService.Insert(accountCreateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *accountController) Update(context *gin.Context) {
	var accountUpdateDTO dto.AccountUpdateDTO
	errDTO := context.ShouldBind(&accountUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.accountService.IsAllowedToEdit(userID, accountUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			accountUpdateDTO.UserID = id
		}
		result := c.accountService.Update(accountUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *accountController) Delete(context *gin.Context) {
	var account entity.Account
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	account.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.accountService.IsAllowedToEdit(userID, account.ID) {
		c.accountService.Delete(account)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *accountController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
