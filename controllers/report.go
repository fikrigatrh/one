package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"number1/auth"
	"number1/models"
	"number1/models/contract"
	"number1/usecase"
	"strconv"
)

type ReportController struct {
	uc usecase.TransReportUsecaseInterface
}

func NewReportControllerImpl(r *gin.RouterGroup, uc usecase.TransReportUsecaseInterface) {
	handler := ReportController{uc}

	r.POST("/report", handler.ReportMerchant)
}

func (r ReportController) ReportMerchant(c *gin.Context) {
	var user models.ReqReportMerchant

	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("page"))
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.Error(errors.New(contract.ErrBadRequest))
		c.Abort()
		return
	}

	tokenAuth, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.Error(errors.New(contract.ErrUnauthorized))
		c.Abort()
		return
	}

	atoi, err := strconv.Atoi(tokenAuth.UserID)
	if err != nil {
		return
	}
	uri := c.Request.Host + c.Request.URL.String()

	report, err := r.uc.GetReportMerchant(user, uri, atoi, limit, offset)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	response := models.ResponseCustom{
		ResponseCode:    "200",
		ResponseMessage: "Successfully",
		Data:            report,
	}
	c.JSON(http.StatusOK, response)
}
