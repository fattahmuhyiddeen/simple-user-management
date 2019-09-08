package model

import (
	"log"
	"time"

	"github.com/fattahmuhyiddeen/simple-user-management/controller/response"
	"github.com/fattahmuhyiddeen/simple-user-management/service"
)

//User is model of user branch
type User struct {
	ID        int       `json:"id" form:"id"`
	Email     string    `json:"email" form:"email"`
	Name      string    `json:"name" form:"name"`
	Token     string    `json:"token" form:"token"`
	Password  string    `json:"password,omitempty" form:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const userTable = "users"

//GetUserByEmail is to search user by email
func GetUserByEmail(email string) (user User) {
	row := db.QueryRow("SELECT id, email, name, password, created_at FROM "+userTable+" WHERE email=$1", email)
	row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt)
	return
}

//GetUserByID is
func GetUserByID(id int) (user User) {
	row := db.QueryRow("SELECT id, email, name, password, created_at FROM "+userTable+" WHERE id=$1", id)
	row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt)

	return
}

//InsertUser is
func InsertUser(user *User) {
	user.UpdatedAt = DateTimeNow()
	user.CreatedAt = user.UpdatedAt

	err := db.QueryRow(
		"INSERT INTO "+userTable+" (email, name, password) VALUES ("+tablePlaceholder("email, name, created_at")+") RETURNING id",
		user.Email, user.Name, user.Password,
	).Scan(&user.ID)
	log.Println(err)
}

func AllUsers() (users []User) {
	users = []User{}
	rows, err := db.Query("SELECT id, name, email, created_at FROM " + userTable + " ORDER BY id")
	if err == nil {
		for rows.Next() {
			temp := new(User)
			rows.Scan(
				&temp.ID,
				&temp.Name,
				&temp.Email,
				&temp.CreatedAt,
			)
			users = append(users, *temp)
		}
	}
	return
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
}
