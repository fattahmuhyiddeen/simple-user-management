package service

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	valid "github.com/asaskevich/govalidator"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	config "github.com/fattahmuhyiddeen/simple-user-management/config"
)

//GenerateToken is to generate JWT token
func GenerateToken(userID string) (result string) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	//sorry, lazy to make refresh token, haha, so set expiry too long. please dont try this at home
	claims["exp"] = time.Now().Add(time.Hour * 99999).Unix()

	// Generate encoded token and send it as response
	result, _ = token.SignedString([]byte(config.GenerateTokenKey))
	return result
}

// IsValidEmail email format
func IsValidEmail(email string) bool {
	return valid.IsEmail(email)
}

// IsValidPassword email format
func IsValidPassword(pwd string) bool {
	if pwd == "" {
		return false
	}
	return true
}

// HashPassword to hash password sent from api
func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		log.Println(err)
		return "error"
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePasswords to check whether password is same
func ComparePasswords(hashedPwd string, pwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.

	return b
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes := GenerateRandomBytes(n)

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

// GenerateRandomStringURLSafe returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomStringURLSafe(n int) string {
	b := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b)
}
