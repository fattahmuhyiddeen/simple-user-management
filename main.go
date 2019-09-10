package main

import (
	"net/http"
	"time"

	config "github.com/fattahmuhyiddeen/simple-user-management/config"
	customMiddleware "github.com/fattahmuhyiddeen/simple-user-management/middleware"
	_ "github.com/fattahmuhyiddeen/simple-user-management/model"
	"github.com/fattahmuhyiddeen/simple-user-management/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Simple User Management
// @version 1.0
// @description This is a sample user management service.

// @contact.name Fattah Muhyiddeen
// @contact.url https://github.com/fattahmuhyiddeen
// @contact.email fattahmuhyiddeen@gmail.com

// @host https://simple-user-management.herokuapp.com/

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Use(middleware.Logger())
	e.Use(customMiddleware.NoCache())

	if config.Env == "heroku" {
		e.Pre(middleware.HTTPSRedirect())
	}
	e.Any("/*", func(c echo.Context) (err error) {
		routes.APIRoutes().ServeHTTP(c.Response(), c.Request())
		return
	})
	serverConfig := &http.Server{
		Addr:         ":" + config.APPPort,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(serverConfig))
}
