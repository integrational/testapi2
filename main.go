package main

import (
	"github.com/integrational/apitests/testapi2/api/gen"
	"github.com/integrational/apitests/testapi2/api/impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	impl := impl.New()
	echo := echo.New()

	echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip} ${method} ${uri} ${status}\n",
		Output: log.Writer(),
	}))
	echo.Use(middleware.Recover())

	gen.RegisterHandlers(echo, impl)

	echo.Logger.Fatal(echo.Start(":" + port))
}
