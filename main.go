package main

import (
	"github.com/victorfernandesraton/psvalidator/command"
	"github.com/victorfernandesraton/psvalidator/infra"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e := echo.New()
	verifyCommand := &command.VerifyPasswordCommand{}
	verifyHttpController := infra.VerifyHttpController{
		VerifyCommand: verifyCommand,
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Runner")
	})

	e.POST("/verify", verifyHttpController.Handler)

	e.Logger.Fatal(e.Start(":3000"))
}
