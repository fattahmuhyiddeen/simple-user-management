package controller

import (
	"net/http"

	model "github.com/fattahmuhyiddeen/simple-user-management/model"

	"github.com/labstack/echo"
)

func AllUsers(c echo.Context) (err error) {
	users := model.AllUsers()
	return c.JSON(http.StatusCreated, users)
}
