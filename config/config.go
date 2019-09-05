package config

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// var DBUser = os.Getenv("DB_USER")
// var DBPassword = os.Getenv("DB_PASSWORD")
// var DBDatabase = os.Getenv("DB_DATABASE")
// var DBDriver = os.Getenv("DB_DRIVER")

// var DBHost = os.Getenv("DB_HOST")
// var DBPort = os.Getenv("DB_PORT")

var GenerateTokenKey = os.Getenv("GENERATE_TOKEN_KEY")
var APPPort = os.Getenv("PORT")
var APPURL = os.Getenv("APP_URL")
var DatabaseURL = os.Getenv("DATABASE_URL")
var SparkPostAPIKey = os.Getenv("SPARKPOST_API_KEY")
var Env = os.Getenv("ENV")
var AESKey = os.Getenv("AES_KEY")

//EmailSender is email used as sender for any automated emails
var EmailSender = "notifier@emeltrack.com"

//google sign in config
const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var googleRedirectURL = os.Getenv("GOOGLE_OAUTH_CALLBACK")
var googleOauthClientID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
var googleOauthClientSecret = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  googleRedirectURL,
	ClientID:     googleOauthClientID,
	ClientSecret: googleOauthClientSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	// Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint: google.Endpoint,
}

const DateTimeFormat = "2006-01-02T15:04:05"
const LogoName = "logo.png"

func init() {
	// if DBUser == "" {
	// 	log.Fatal("cannot find DB_USER from env")
	// }

	// if DBPassword == "" {
	// 	log.Fatal("cannot find DB_PASSWORD from env")
	// }

	// if DBHost == "" {
	// 	log.Fatal("cannot find DB_HOST from env")
	// }

	// if DBPort == "" {
	// 	log.Fatal("cannot find DB_PORT from env")
	// }

	// if DBDatabase == "" {
	// 	log.Fatal("cannot find DB_DATABASE from env")
	// }

	if GenerateTokenKey == "" {
		log.Fatal("cannot find GENERATE_TOKEN_KEY from env")
	}

	if AESKey == "" {
		log.Fatal("cannot find AES_KEY from env")
	}

	// if APPPort == "" {
	// 	log.Fatal("cannot find PORT from env")
	// }
	// if APPURL == "" {
	// 	log.Fatal("cannot find APP_URL from env")
	// }

	log.Println(DatabaseURL)
}
