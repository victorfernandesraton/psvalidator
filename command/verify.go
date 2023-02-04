package command

import (
	"errors"
	"psvalidator/domain"
)

var NotValidRuleError = errors.New("not valid rule")

type VerifyPasswordCommand struct {
}

type VerifyPasswordCommandParams struct {
	Password string
	Rules    []*domain.Rule
}

type VerifyPasswordCommandResponse struct {
	Verify  bool
	NoMatch []domain.RuleEnum
}

func (c *VerifyPasswordCommand) Execute(params VerifyPasswordCommandParams) (*VerifyPasswordCommandResponse, error) {
	var noMatchRules []domain.RuleEnum
	for _, rule := range params.Rules {
		var err error

		switch rule.Type {
		case domain.SizeRule:
			err = domain.MinSize(params.Password, rule.Value)
		case domain.LowerRule:
			err = domain.MinLowerCase(params.Password, rule.Value)
		case domain.UpperRule:
			err = domain.MinUpperCase(params.Password, rule.Value)
		case domain.DigitRule:
			err = domain.MinDigit(params.Password, rule.Value)
		case domain.SpecialRule:
			err = domain.MinSpecialChars(params.Password, rule.Value)
		case domain.RepeatRule:
			err = domain.NoRepeat(params.Password)
		default:
			return nil, NotValidRuleError
		}
		if err != nil {
			var errorInRules bool
			for _, errorForRule := range domain.RuleErrors {
				if errorForRule == err {
					errorInRules = true
				}
			}
			if !errorInRules {
				return nil, err
			}
			noMatchRules = append(noMatchRules, rule.Type)
		}
	}

	return &VerifyPasswordCommandResponse{
		Verify:  len(noMatchRules) == 0,
		NoMatch: noMatchRules,
	}, nil
}
