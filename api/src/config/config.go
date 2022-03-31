package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	SecretKey          []byte
)

func Carregar() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Porta = 5000
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True", user, password, db_name)

	SecretKey = []byte(os.Getenv("SECRET_ENV"))

}
