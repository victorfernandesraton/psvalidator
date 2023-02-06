package infra

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorfernandesraton/psvalidator/command"
	"github.com/victorfernandesraton/psvalidator/domain"
)

type VerifyCommandInterface interface {
	Execute(params command.VerifyPasswordCommandParams) (*command.VerifyPasswordCommandResponse, error)
}

type VerifyHttpController struct {
	VerifyCommand VerifyCommandInterface
}

type VerifyBody struct {
	Password string         `json:"password"`
	Rules    []*domain.Rule `json:"rules"`
}

func (ctr *VerifyHttpController) Handler(c echo.Context) error {

	body := new(VerifyBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := ctr.VerifyCommand.Execute(command.VerifyPasswordCommandParams{
		Password: body.Password,
		Rules:    body.Rules,
	})

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == command.NotValidRuleError {
			statusCode = http.StatusBadRequest
		}
		return echo.NewHTTPError(statusCode, err)
	}

	return c.JSON(http.StatusOK, res)
}
