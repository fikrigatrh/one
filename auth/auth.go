package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"number1/models"
	"os"
	"strings"
	"time"
)

func CreateToken(authD models.Auth) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_uuid"] = authD.AuthUUID
	claims["username"] = authD.Username
	claims["user_id"] = authD.UserID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request) (*models.Auth, error) {
	var res models.Auth
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authUuid, ok := claims["auth_uuid"].(string)
		if !ok {
			return nil, err
		}
		username, done := claims["username"].(string)
		if !done {
			return nil, err
		}
		userId, done := claims["user_id"].(string)
		//userIdRes, err := strconv.Atoi(userId)
		//if err != nil {
		//	return nil, err
		//}
		if !done {
			return nil, err
		}

		res.AuthUUID = authUuid
		res.Username = username
		res.UserID = userId
		return &res, nil

	}
	return nil, err
}
