package autenticacao

import (
	"devbook-api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriaToken(usuarioId int64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	signedString, _ := token.SignedString([]byte(config.SecretKey))

	return signedString, nil
}
