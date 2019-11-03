package main

import (
	"github.com/integrational/apitests/testapi2/api/gen"
	"github.com/integrational/apitests/testapi2/api/impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	gen.RegisterHandlers(echo, impl.New())

	echo.Logger.Fatal(echo.Start(":" + port))
}
