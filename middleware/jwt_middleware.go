package middleware

import (
	"ci_cd_golang/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 240).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.SECRET_JWT))
}
