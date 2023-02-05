package command_test

import (
	"github.com/victorfernandesraton/psvalidator/command"
	"github.com/victorfernandesraton/psvalidator/domain"
	"testing"
)

var stub = &command.VerifyPasswordCommand{}

func TestVerifyPasswordCommand_Execute(t *testing.T) {
	var cases = []struct {
		desc   string
		params command.VerifyPasswordCommandParams
		res    *command.VerifyPasswordCommandResponse
		err    error
	}{
		{
			desc: "Validate example for digit error",
			params: command.VerifyPasswordCommandParams{
				Password: "TesteSenhaForte!123&",
				Rules: []*domain.Rule{
					{
						Type:  domain.SizeRule,
						Value: 8,
					},
					{
						Type:  domain.SpecialRule,
						Value: 2,
					},
					{
						Type:  domain.RepeatRule,
						Value: 0,
					},
					{
						Type:  domain.DigitRule,
						Value: 4,
					},
				},
			},
			res: &command.VerifyPasswordCommandResponse{
				Verify:  false,
				NoMatch: []domain.RuleEnum{domain.DigitRule},
			},
		},
		{
			desc: "Validate example with no rule",
			params: command.VerifyPasswordCommandParams{
				Password: "TesteSenhaForte!123&",
			},
			res: &command.VerifyPasswordCommandResponse{
				Verify: true,
			},
		},
		{
			desc: "Validate whitout password",
			params: command.VerifyPasswordCommandParams{
				Password: "",
				Rules: []*domain.Rule{
					{
						Type:  domain.SizeRule,
						Value: 8,
					},
					{
						Type:  domain.SpecialRule,
						Value: 2,
					},
					{
						Type:  domain.RepeatRule,
						Value: 0,
					},
					{
						Type:  domain.DigitRule,
						Value: 4,
					},
				},
			},
			res: &command.VerifyPasswordCommandResponse{
				Verify:  false,
				NoMatch: []domain.RuleEnum{domain.DigitRule, domain.SpecialRule, domain.SizeRule},
			},
		},
		{
			desc: "Validate with repetition password",
			params: command.VerifyPasswordCommandParams{
				Password: "testeSeenhaForte!123&",
				Rules: []*domain.Rule{
					{
						Type:  domain.SizeRule,
						Value: 8,
					},
					{
						Type:  domain.SpecialRule,
						Value: 2,
					},
					{
						Type:  domain.RepeatRule,
						Value: 0,
					},
					{
						Type:  domain.DigitRule,
						Value: 4,
					},
				},
			},
			res: &command.VerifyPasswordCommandResponse{
				Verify:  false,
				NoMatch: []domain.RuleEnum{domain.DigitRule, domain.RepeatRule},
			},
		},
	}
	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			res, err := stub.Execute(t1.params)
			if err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
			if res.Verify != t1.res.Verify {
				t.Fatalf("expect %v , got %v", t1.res.Verify, res.Verify)
			}
			if len(t1.res.NoMatch) != len(res.NoMatch) {
				t.Fatalf("expect size list no match%v , got %v", t1.res.NoMatch, res.NoMatch)
			}
		})
	}
}
