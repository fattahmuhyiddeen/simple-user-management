package main

import (
	"net/http"
	"time"

	config "github.com/fattahmuhyiddeen/simple-user-management/config"
	customMiddleware "github.com/fattahmuhyiddeen/simple-user-management/middleware"
	"github.com/fattahmuhyiddeen/simple-user-management/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

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
