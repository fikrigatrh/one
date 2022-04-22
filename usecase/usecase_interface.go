package usecase

import "number1/models"

type LoginUsecaseInterface interface {
	CreateAuth(user models.User) (*models.Auth, error)
	SignIn(authD models.Auth) (string, error)
	DeleteAuth(uuid string) error
}

type TransReportUsecaseInterface interface {
	GetReportMerchant(req models.ReqReportMerchant, uri string, id, limit, offset int) (models.ResponseDataPagination, error)
}

type ErrorHandlerUsecase interface {
	ResponseError(error interface{}) (int, interface{})
	ValidateRequest(error interface{}) (string, error)
}
