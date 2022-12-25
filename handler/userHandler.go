package handler

import (
	"MALIKI-KARIM/dto"
	"MALIKI-KARIM/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userRestHandler struct {
	service service.UserService
}

func newUserHandler(userService service.UserService) userRestHandler {
	return userRestHandler{
		service: userService,
	}
}

func (u userRestHandler) Register(c *gin.Context) {
	var user dto.RegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": err.Error(),
		})
		return
	}

	result, err := u.service.Register(&user)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, dto.DataRegisterResponse(*result))
}

func (u userRestHandler) Login(c *gin.Context) {
	var user dto.LoginRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.Login(&user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": err.Message(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u userRestHandler) Deposit(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}
	var depositRequest dto.DepositRequest
	if err := c.ShouldBindJSON(&depositRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.Deposit(userId, &depositRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": http.StatusText(http.StatusInternalServerError),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u userRestHandler) Transfer(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	var transferRequest dto.TransferRequest
	if err := c.ShouldBindJSON(&transferRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}
	result, err := u.service.Transfer(userId, &transferRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": http.StatusText(http.StatusInternalServerError),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
