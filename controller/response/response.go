package response

import (
	"net/http"
	"github.com/labstack/echo"
)

func BadRequest(message string)(err error) {
	return &echo.HTTPError{Code: http.StatusBadRequest, Message: message}
}