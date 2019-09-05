package service

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

//SetCookie is
func SetCookie(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	c.SetCookie(cookie)
}
