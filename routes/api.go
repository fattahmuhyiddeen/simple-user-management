package routes

import (
	"strings"

	config "github.com/fattahmuhyiddeen/simple-user-management/config"
	controller "github.com/fattahmuhyiddeen/simple-user-management/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// APIRoutes :
func APIRoutes() *echo.Echo {
	api := echo.New()
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	// api.Logger.SetLevel(log.ERROR)
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.GenerateTokenKey),
		Skipper: func(c echo.Context) bool {

			if strings.Contains(c.Path(), "/public/") {
				return true
			}

			switch c.Path() {
			case "/swagger/*", "/login", "/register", "/ping":
				return true
			default:
				return false
			}
		},
	}))

	// Routes
	api.GET("/", controller.HomePage)
	api.GET("/ping", controller.Ping)

	return api
}
