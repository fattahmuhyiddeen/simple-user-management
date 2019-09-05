package model

import (
	"log"

	"github.com/fattahmuhyiddeen/simple-user-management/controller/response"
	"github.com/fattahmuhyiddeen/simple-user-management/service"
)

// config "github.com/fattahmuhyiddeen/simple-user-management/config"

//User is model of user branch
type User struct {
	ID        int    `json:"id" form:"id"`
	Email     string `json:"email" form:"email"`
	Name      string `json:"name" form:"name"`
	Picture   string `json:"picture" form:"picture"`
	Token     string `json:"token" form:"token"`
	Password  string `json:"password,omitempty" form:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const userTable = "users"
const userFields = "email, name, password, picture, created_at, updated_at"

//GetUserByEmail is to search user by email
func GetUserByEmail(email string) (user User) {
	connectDB()
	defer disconnectDB()

	row := db.QueryRow("SELECT id, "+userFields+" FROM "+userTable+" WHERE email=$1", email)
	row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Picture, &user.CreatedAt, &user.UpdatedAt)
	return
}

//GetUserByID is
func GetUserByID(id int) (user User) {
	connectDB()
	defer disconnectDB()

	row := db.QueryRow("SELECT id, "+userFields+" FROM "+userTable+" WHERE id=$1", id)
	row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Picture, &user.CreatedAt, &user.UpdatedAt)

	return
}

//InsertUser is
func InsertUser(user *User) {
	connectDB()
	defer disconnectDB()

	user.UpdatedAt = DateTimeNow()
	user.CreatedAt = user.UpdatedAt

	err := db.QueryRow(
		"INSERT INTO "+userTable+" ("+userFields+") VALUES ("+tablePlaceholder(userFields)+") RETURNING id",
		user.Email, user.Name, user.Password, &user.Picture, user.CreatedAt, user.UpdatedAt,
	).Scan(&user.ID)
	log.Println(err)
}

//------------------------------------------------------
// Function below dont connect to DB
//------------------------------------------------------

//IsEmailExist to check if email already exist in user table
func IsEmailExist(email string) bool {
	return GetUserByEmail(email).ID != 0
}

//ValidateUser to
func ValidateUser(user *User) (err error) {
	if !service.IsValidEmail(user.Email) {
		return response.BadRequest("invalid email")
	}
	if !service.IsValidPassword(user.Password) {
		return response.BadRequest("invalid password")
	}

	if user.Name == "" {
		return response.BadRequest("invalid name")
	}

	if IsEmailExist(user.Email) {
		return response.BadRequest(user.Email + " already taken")
	}
	return
}

func ClearUserSensitiveFields(user *User) {
	user.Password = ""
	user.CreatedAt = ""
	user.UpdatedAt = ""
}
