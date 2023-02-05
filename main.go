package main

import (
	"github.com/victorfernandesraton/psvalidator/command"
	"github.com/victorfernandesraton/psvalidator/infra"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func ServerFactory() *echo.Echo {

	e := echo.New()
	verifyCommand := &command.VerifyPasswordCommand{}
	verifyHttpController := infra.VerifyHttpController{
		VerifyCommand: verifyCommand,
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Runner")
	})

	e.POST("/verify", verifyHttpController.Handler)

	return e
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	server := ServerFactory()
	server.Logger.Fatal(server.Start(":" + port))
}
