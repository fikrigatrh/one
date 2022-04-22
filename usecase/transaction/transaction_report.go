package transaction

import (
	"errors"
	"number1/config/env"
	"number1/models"
	"number1/models/contract"
	"number1/repo"
	"number1/usecase"
	"strconv"
	"strings"
)

type TransReportUsecase struct {
	repo repo.TransReportRepoInterface
}

func NewTransReportUsecase(repo repo.TransReportRepoInterface) usecase.TransReportUsecaseInterface {
	return TransReportUsecase{repo: repo}
}

func (t TransReportUsecase) GetReportMerchant(req models.ReqReportMerchant, uri string, id, limit, offset int) (models.ResponseDataPagination, error) {
	var omzet int
	if limit == 0 {
		limit = 1
	}

	if offset <= 0 {
		offset = 1
	}

	if req.MerchantId != "" {
		err := t.GetIdMerchant(id, req.MerchantId)
		if err != nil {
			return models.ResponseDataPagination{}, err
		}
	} else {
		err := t.GetIdOutlet(id, req.OutletId)
		if err != nil {
			return models.ResponseDataPagination{}, err
		}
	}

	report, tempArr, err := t.repo.GetReport(req, limit, offset)
	if err != nil {
		return models.ResponseDataPagination{}, err
	}

	lengthData := report.TotalData
	eachData := report.NumberEnd

	if lengthData == 0 {
		lengthData = limit
	}

	if limit > lengthData {
		limit = lengthData
	}

	tempTotalPage := lengthData / limit

	if lengthData%limit != 0 {
		tempTotalPage = tempTotalPage + 1
	}

	report.TotalData = lengthData
	report.DataPerPage = limit
	report.TotalPage = tempTotalPage

	if eachData < limit {
		report.NumberEnd = (limit * (offset - 1)) + eachData
	} else {
		report.NumberEnd = limit * offset
	}

	if offset == 1 {
		report.NumberCurrent = 1
	} else if report.TotalPage == offset && eachData == 1 {
		report.NumberCurrent = report.NumberEnd
	} else if report.TotalPage == offset && eachData != 1 {
		report.NumberCurrent = (report.NumberEnd - eachData) + 1
	} else {
		report.NumberCurrent = (report.NumberEnd - limit) + 1
	}

	report.Page = offset

	for i2, _ := range report.Data {
		report.Data[i2].No = i2 + report.NumberCurrent
	}

	tempUrl := strings.Split(uri, "&page=")
	atoi, _ := strconv.Atoi(tempUrl[1])
	tempPageNext := atoi + 1
	tempPagePrev := atoi - 1
	if tempPagePrev < 1 {
		tempPagePrev = 1
	}
	nextPage := strconv.Itoa(tempPageNext)
	prevPage := strconv.Itoa(tempPagePrev)

	for _, m := range tempArr {
		omzet += m.BillTotal
	}

	report.Omzet = omzet

	if report.Page <= 1 {
		if report.DataPerPage == report.TotalData {
			report.PrevUrlPage = ""
			report.NextUrlPage = ""
		} else {
			report.PrevUrlPage = ""
			report.NextUrlPage = env.Config.Protocol + "://" + tempUrl[0] + "&page=" + nextPage
		}
	} else if report.Page == report.TotalPage {
		report.PrevUrlPage = env.Config.Protocol + "://" + tempUrl[0] + "&page=" + prevPage
		report.NextUrlPage = ""
	} else {
		report.PrevUrlPage = env.Config.Protocol + "://" + tempUrl[0] + "&page=" + prevPage
		report.NextUrlPage = env.Config.Protocol + "://" + tempUrl[0] + "&page=" + nextPage
	}

	if report.Data == nil {
		return models.ResponseDataPagination{}, errors.New(contract.ErrDataNotFound)

	}
	return report, nil
}

func (t TransReportUsecase) GetIdMerchant(id int, merchantId string) error {
	merchant, err := t.repo.GetIdMerchant(id)
	if err != nil {
		return err
	}

	mId, _ := strconv.Atoi(merchantId)
	idRes := uint(mId)
	for _, m := range merchant {
		if m.ID != idRes {
			continue
		} else {
			return nil
		}
	}

	return errors.New(contract.ErrMerchantNotFound)
}

func (t TransReportUsecase) GetIdOutlet(id int, outletId string) error {

	merchant, err := t.repo.GetIdMerchant(id)
	if err != nil {
		return err
	}

	mId, _ := strconv.Atoi(outletId)
	idRes := uint(mId)
	for _, m := range merchant {
		outlet, err := t.repo.GetIdOutlet(int(m.ID))
		if err != nil {
			return err
		}
		for _, m2 := range outlet {
			if m2.ID != idRes {
				continue
			} else {
				return nil
			}
		}
	}

	return errors.New(contract.ErrOutletNotFound)
}
