package repo

import "number1/models"

type LoginRepoInterface interface {
	CheckUser(v models.User) (models.User, error)
	CreateAuth(username string, userId int) (*models.Auth, error)
	DeleteAuth(uuid string) error
}

type TransReportRepoInterface interface {
	GetReport(req models.ReqReportMerchant, limit, offset int) (models.ResponseDataPagination, []models.Report, error)
	GetIdMerchant(id int) ([]models.Merchant, error)
	GetIdOutlet(id int) ([]models.Outlet, error)
}
