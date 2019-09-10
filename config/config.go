package config

import (
	"log"
	"os"
)

var GenerateTokenKey = os.Getenv("GENERATE_TOKEN_KEY")

// APPPort is to get PORT from env. In heroku, PORT is declared implicitly and randomly
var APPPort = os.Getenv("PORT")
var DatabaseURL = os.Getenv("DATABASE_URL")
var Env = os.Getenv("ENV")
var AESKey = os.Getenv("AES_KEY")

const DateTimeFormat = "2006-01-02T15:04:05"

func init() {
	if GenerateTokenKey == "" {
		log.Fatal("cannot find GENERATE_TOKEN_KEY from env")
	}

	if AESKey == "" {
		log.Fatal("cannot find AES_KEY from env")
	}
}
