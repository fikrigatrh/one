package middlewares

import (
	"github.com/gin-gonic/gin"
	"number1/usecase"
)

type ErrorHandler struct {
	ErrorHandlerUsecase usecase.ErrorHandlerUsecase
}

func NewErrorHandler(r *gin.RouterGroup, ehus usecase.ErrorHandlerUsecase) {
	handler := &ErrorHandler{
		ErrorHandlerUsecase: ehus,
	}

	r.Use(handler.errorHandler)
}

func (e *ErrorHandler) errorHandler(c *gin.Context) {
	c.Next()

	errorToPrint := c.Errors.Last()
	if errorToPrint != nil {
		c.JSON(e.ErrorHandlerUsecase.ResponseError(errorToPrint))
		c.Abort()
		return
	}
}
