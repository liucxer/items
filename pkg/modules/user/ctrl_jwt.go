package user

import (
	"time"

	"git.querycap.com/ss/srv-aisys/constants/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/saitofun/qlib/encoding"
)

type JWTClaims struct {
	Username string
	Password string
	jwt.StandardClaims
}

var key = encoding.StrToBytes("npFgbpGhLs4cf5e7m1YaemXTYOjFgtkkk")

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := JWTClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "smart-station",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func ParseToken(val string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		val,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)
	if err != nil {
		return nil, errors.UserLoginInvalidToken
	}

	if token != nil {
		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.UserLoginInvalidToken
}
