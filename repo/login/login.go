package login

import (
	"errors"
	"github.com/twinj/uuid"
	"gorm.io/gorm"
	"number1/models"
	"number1/models/contract"
	"number1/repo"
	"strconv"
)

type LoginRepoStruct struct {
	db *gorm.DB
}

func NewLoginRepoImpl(db *gorm.DB) repo.LoginRepoInterface {
	return LoginRepoStruct{db: db}
}

func (l LoginRepoStruct) CheckUser(v models.User) (models.User, error) {
	err := l.db.Debug().Where("user_name = ? and password = ?", v.Username, v.Password).First(&v).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, errors.New(contract.ErrUserNotFound)
		}
		return models.User{}, err
	}

	return v, nil
}

func (l LoginRepoStruct) CreateAuth(username string, userId int) (*models.Auth, error) {
	au := &models.Auth{}
	tx := l.db.Begin()

	id := strconv.Itoa(userId)

	au.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	au.Username = username
	au.UserID = id
	err := l.db.Debug().Create(&au).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return au, nil
}

func (l LoginRepoStruct) DeleteAuth(uuid string) error {
	au := models.Auth{}
	err := l.db.Debug().Where("auth_uuid = ?", uuid).Delete(&au).Error
	if err != nil {
		return err
	}

	return nil
}
