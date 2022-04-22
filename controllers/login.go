package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"number1/auth"
	"number1/models"
	"number1/models/contract"
	"number1/usecase"
)

type LoginController struct {
	LoginUsecase usecase.LoginUsecaseInterface
	err          usecase.ErrorHandlerUsecase
}

func NewLoginControllerImpl(r *gin.RouterGroup, LoginUsecase usecase.LoginUsecaseInterface, err usecase.ErrorHandlerUsecase) {
	handler := LoginController{LoginUsecase, err}
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)
}

func (l LoginController) Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.Error(errors.New(contract.ErrBadRequest))
		c.Abort()
		return
	}

	var authD models.Auth

	fieldErr, err := l.err.ValidateRequest(user)
	if err != nil {
		c.Error(err).SetMeta(models.ErrMeta{
			FieldErr: fieldErr,
		})
		c.Abort()
		return
	}

	authRes, err := l.LoginUsecase.CreateAuth(user)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	authD.AuthUUID = authRes.AuthUUID
	authD.Username = authRes.Username
	authD.UserID = authRes.UserID

	var jwt models.JwtModel

	in, err := l.LoginUsecase.SignIn(authD)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	jwt.Token = in

	response := models.ResponseCustom{
		ResponseCode:    "200",
		ResponseMessage: "Successfully",
		Data:            jwt,
	}
	c.JSON(http.StatusOK, response)

}

func (l LoginController) Logout(c *gin.Context) {
	tokenAuth, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.Error(errors.New(contract.ErrUnauthorized))
		c.Abort()
		return
	}

	err = l.LoginUsecase.DeleteAuth(tokenAuth.AuthUUID)
	if err != nil {
		c.Error(errors.New(contract.ErrUnauthorized))
		c.Abort()
		return
	}

	response := models.ResponseCustom{
		ResponseCode:    "200",
		ResponseMessage: "Successfully",
	}
	c.JSON(http.StatusOK, response)
}
