package autenticacao

import (
	"devbook-api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDg3NDU1MzYsInVzdWFyaW9JZCI6Mn0.KqYX3wvY9L3CXvP8Sd1BmCfUNvOSo41OfXAAtECi1dw

func CriaToken(usuarioId int64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	signedString, _ := token.SignedString([]byte(config.SecretKey))

	return signedString, nil
}

func ValidarToken(r *http.Request) error {
	tokenString := extraiToken(r)
	token, err := jwt.Parse(tokenString, retornaChaveDeVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalido")
}

func extraiToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornaChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperado. %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
