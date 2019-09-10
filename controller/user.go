package controller

import (
	"net/http"

	model "github.com/fattahmuhyiddeen/simple-user-management/model"

	"github.com/labstack/echo"
)

// AllUsers godoc
// @Summary Get list of all users
// @Description Get list of all users
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Router /all_users [get]
func AllUsers(c echo.Context) (err error) {
	users := model.AllUsers()
	return c.JSON(http.StatusOK, users)
}
