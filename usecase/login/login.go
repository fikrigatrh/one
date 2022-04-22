package login

import (
	"crypto/md5"
	"encoding/hex"
	"number1/auth"
	"number1/models"
	"number1/repo"
	"number1/usecase"
)

type LoginUsecaseStruct struct {
	LoginRepo repo.LoginRepoInterface
}

func NewLoginUsecaseImpl(LoginRepo repo.LoginRepoInterface) usecase.LoginUsecaseInterface {
	return &LoginUsecaseStruct{LoginRepo}
}

func (l LoginUsecaseStruct) CreateAuth(user models.User) (*models.Auth, error) {
	hasher := md5.New()
	hasher.Write([]byte(user.Password))
	pass := hex.EncodeToString(hasher.Sum(nil))

	user.Password = pass

	checkUser, err := l.LoginRepo.CheckUser(user)
	if err != nil {
		return nil, err
	}

	auth, err := l.LoginRepo.CreateAuth(checkUser.Username, int(checkUser.ID))
	if err != nil {
		return nil, err
	}

	auth.Username = checkUser.Username
	return auth, nil
}

func (l LoginUsecaseStruct) SignIn(authD models.Auth) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (l LoginUsecaseStruct) DeleteAuth(uuid string) error {
	return l.LoginRepo.DeleteAuth(uuid)
}
