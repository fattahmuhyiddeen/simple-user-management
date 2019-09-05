package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/fattahmuhyiddeen/simple-user-management/model"

	"github.com/labstack/echo"
)

//Ping is used to healthcheck only
func Ping(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "Pong")
}

func HomePage(c echo.Context) (err error) {
	var results map[string]interface{}
	json.Unmarshal([]byte(`{ "version" : "0.1.3"}`), &results)
	return c.JSON(http.StatusOK, results)
}

func CheckDB(c echo.Context) (err error) {
	var results map[string]interface{}
	var status = "failed"
	if model.TestConnection() {
		status = "success"
	}
	json.Unmarshal([]byte(`{ "connection to DB" : "`+status+`"}`), &results)
	return c.JSON(http.StatusOK, results)
}
