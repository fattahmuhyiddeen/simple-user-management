package controller

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	response "github.com/fattahmuhyiddeen/simple-user-management/controller/response"
	model "github.com/fattahmuhyiddeen/simple-user-management/model"
	service "github.com/fattahmuhyiddeen/simple-user-management/service"
	"github.com/labstack/echo"
)

//UserStruct is consist of User model, request and response
type UserStruct struct {
	model.User
}

//Register is
func Register(c echo.Context) (err error) {
	user := new(UserStruct)
	c.Bind(user)

	// Validate
	if err = model.ValidateUser(&user.User); err != nil {
		return
	}

	user.Password = service.HashPassword(user.Password)
	user.Token = service.GenerateToken(strconv.Itoa(user.ID))
	model.InsertUser(&user.User)

	//remove unnecessary variables from response
	model.ClearUserSensitiveFields(&user.User)
	return c.JSON(http.StatusCreated, user)
}

//Login is
func Login(c echo.Context) (err error) {
	user := new(UserStruct)
	c.Bind(user)

	searchUser := model.GetUserByEmail(user.Email)

	if searchUser.ID == 0 || !service.ComparePasswords(searchUser.Password, user.Password) {
		return response.BadRequest("Login credential not match")
	}
	user.User = searchUser
	user.Token = service.GenerateToken(strconv.Itoa(user.ID))

	//remove unnecessary variables from response
	model.ClearUserSensitiveFields(&user.User)
	return c.JSON(http.StatusCreated, user)
}

func userIDFromToken(c echo.Context) (id int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id, _ = strconv.Atoi(claims["id"].(string))
	return
}

func getAuthUser(c echo.Context) model.User {
	return model.GetUserByID(userIDFromToken(c))
}
