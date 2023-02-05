package infra

import (
	"github.com/labstack/echo/v4"
	"github.com/victorfernandesraton/psvalidator/command"
	"github.com/victorfernandesraton/psvalidator/domain"
	"log"
	"net/http"
)

type VerifyCommandInterface interface {
	Execute(params command.VerifyPasswordCommandParams) (*command.VerifyPasswordCommandResponse, error)
}

type VerifyHttpController struct {
	VerifyCommand VerifyCommandInterface
}

type VerifyBodyRule struct {
	Rule  string `json:"rule"`
	Value int    `json:"value"`
}

type VerifyBody struct {
	Password string           `json:"password"`
	Rules    []VerifyBodyRule `json:"rules,omitempty"`
}

func RuleForObject(rule *VerifyBodyRule) *domain.Rule {
	return &domain.Rule{
		Type:  rule.Rule,
		Value: rule.Value,
	}
}

func (ctr *VerifyHttpController) Handler(c echo.Context) error {

	var rules []*domain.Rule
	body := new(VerifyBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	for _, rule := range body.Rules {
		rules = append(rules, RuleForObject(&rule))
	}

	log.Println(rules)

	res, err := ctr.VerifyCommand.Execute(command.VerifyPasswordCommandParams{
		Password: body.Password,
		Rules:    rules,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
