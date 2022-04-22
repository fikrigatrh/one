package error

import (
	"errors"
	"fmt"
	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"net/http"
	"number1/models"
	"number1/models/contract"
	"number1/usecase"
	"number1/utils"
	"strconv"
	"strings"
)

type errorHandlerUsecase struct {
}

func NewErrorHandlerUsecase() usecase.ErrorHandlerUsecase {
	return &errorHandlerUsecase{}
}

func (eh *errorHandlerUsecase) ResponseError(A interface{}) (int, interface{}) {
	var T interface{}
	var fieldNameErr string
	var serviceCode string

	if A.(*gin.Error).Meta != nil {
		fieldNameErr = A.(*gin.Error).Meta.(models.ErrMeta).FieldErr
	}

	// Check A is a correct error type and assign to T
	if A.(*gin.Error).Err != nil {
		T = A.(*gin.Error).Err
	}

	fmt.Println(serviceCode)

	switch T.(type) {
	case error:
		if _, ok := T.(*pq.Error); ok {
			switch T.(*pq.Error).Code.Name() {
			case "unique_violation":
				return errcntrct.ErrorMessage(http.StatusInternalServerError, "", errors.New(contract.ErrGeneralError))
			}
		}

		switch T.(error).Error() {
		case contract.ErrInvalidFieldFormat:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, contract.ErrInvalidFieldFormat, fieldNameErr)
		case contract.ErrInvalidFieldMandatory:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, contract.ErrInvalidFieldMandatory, fieldNameErr)
		case contract.ErrBadRequest:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, "", "")
		case contract.ErrGeneralError:
			return responseErrorAdapter(T.(error), http.StatusInternalServerError, "", "")
		case contract.ErrUnauthorized:
			return responseErrorAdapter(T.(error), http.StatusUnauthorized, "", "")
		case contract.ErrUserNotFound:
			return responseErrorAdapter(T.(error), http.StatusNotFound, "", "")
		case contract.ErrMerchantNotFound:
			return responseErrorAdapter(T.(error), http.StatusNotFound, "", "")
		case contract.ErrOutletNotFound:
			return responseErrorAdapter(T.(error), http.StatusNotFound, "", "")
		case contract.ErrDataNotFound:
			return responseErrorAdapter(T.(error), http.StatusNotFound, "", "")
		default:
			return responseErrorAdapter(errors.New(contract.ErrGeneralError), http.StatusInternalServerError, "", "")
		}
	}

	return responseErrorAdapter(T.(error), http.StatusInternalServerError, "", "")
}

var Case string

func (eh *errorHandlerUsecase) ValidateRequest(T interface{}) (string, error) {
	v := validator.New()
	var errArr error
	var field string
	switch T.(type) {
	case models.User:
		err := v.Struct(T)
		if err == nil {
			return "", nil
		}
		for _, e := range err.(validator.ValidationErrors) {
			if e.Value() != "" {
				switch e.Tag() {
				case "numeric", "max", "email", "lt", "gte", "len", "alpha", "alphanum":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldFormat)
				}
				break
			} else {
				switch e.Tag() {
				case "required":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldMandatory)
				}
				break
			}
		}

		if errArr != nil {
			return field, errArr
		}

		return "", nil

	default:
		return "", errors.New(contract.ErrGeneralError)
	}

}

func responseErrorAdapter(errHttpStatus interface{}, httpStatusCode int, ctr string, fieldErr string) (int, models.ResponseCustomErr) {
	_, errData := errcntrct.ErrorMessage(httpStatusCode, "", errHttpStatus)
	var resp models.ResponseCustomErr

	resp.ResponseCode = strconv.Itoa(httpStatusCode)
	if strings.Contains(contract.FieldErr, " ") {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg, contract.FieldErr)
	} else if ctr == "400001" || ctr == "400002" {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg, utils.LowerCamelCase(fieldErr))
	} else {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg)
	}
	return httpStatusCode, resp
}
