package transaction

import (
	"gorm.io/gorm"
	"number1/models"
	"number1/repo"
	"strings"
)

type TransReportRepo struct {
	db *gorm.DB
}

func NewTransReportRepo(db *gorm.DB) repo.TransReportRepoInterface {
	return &TransReportRepo{db}
}

var (
	queryGetMerchant = `SELECT distinct m.merchant_name, o.outlet_name, t.created_at as date, t.bill_total as bill_total
from transactions t, merchants m, outlets o
where date (t.created_at) = ? and t.merchant_id = ?
GROUP BY t.id
order by t.id desc limit ? offset ?;`
	queryGetOutlet = `SELECT distinct m.merchant_name, o.outlet_name, t.created_at as date, t.bill_total as bill_total
from transactions t, merchants m, outlets o
where date (t.created_at) = ? and t.outlet_id = ?
GROUP BY t.id
order by t.id desc limit ? offset ?;`
)

func (t TransReportRepo) GetReport(req models.ReqReportMerchant, limit, offset int) (models.ResponseDataPagination, []models.Report, error) {

	var pagination models.ResponseDataPagination
	offset = (limit * offset) - limit
	var query string

	var id string
	if req.MerchantId == "" {
		query = queryGetOutlet
		id = req.OutletId
	} else {
		query = queryGetMerchant
		id = req.MerchantId
	}

	var v []models.Report
	err := t.db.Debug().Raw(query, req.Date, id, limit, offset).Scan(&v).Error
	if err != nil {
		return models.ResponseDataPagination{}, nil, err
	}

	var tempReport models.Report
	var tempArr []models.Report
	QueryTotalPage := QueryTotalPage(query)
	rowTempArray, err := t.db.Debug().Raw(QueryTotalPage, req.Date, id).Rows()
	defer rowTempArray.Close()
	for rowTempArray.Next() {
		t.db.ScanRows(rowTempArray, &tempReport)
		if err != nil {
			return models.ResponseDataPagination{}, nil, err
		}
		tempArr = append(tempArr, tempReport)
	}

	length := len(tempArr)
	pagination.TotalData = length
	pagination.NumberEnd = len(v)

	pagination.Data = v

	return pagination, tempArr, nil
}

func (t TransReportRepo) GetIdMerchant(id int) ([]models.Merchant, error) {
	var v []models.Merchant
	err := t.db.Debug().Model(&models.Merchant{}).Where("user_id = ?", id).Scan(&v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (t TransReportRepo) GetIdOutlet(id int) ([]models.Outlet, error) {
	var v []models.Outlet
	err := t.db.Debug().Model(&models.Outlet{}).Where("merchant_id = ?", id).Scan(&v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func QueryTotalPage(query string) string {
	ArrayQuery := strings.Split(query, " ")
	jumlah := len(ArrayQuery) - 4
	tempQuery := ArrayQuery[:jumlah]
	str := strings.Join(tempQuery, " ")
	fixQuery := str + ";"
	return fixQuery
}
