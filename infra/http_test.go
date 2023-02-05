package infra_test

import (
	"github.com/victorfernandesraton/psvalidator/domain"
	"github.com/victorfernandesraton/psvalidator/infra"
	"testing"
)

func TestRuleTransform(t *testing.T) {

	cases := []struct {
		input  *infra.VerifyBodyRule
		output *domain.Rule
	}{
		{
			input: &infra.VerifyBodyRule{
				Rule:  "noRepeat",
				Value: 0,
			},
			output: &domain.Rule{
				Type:  "noRepeat",
				Value: 0,
			},
		},
		{
			input: &infra.VerifyBodyRule{
				Rule:  "minDigit",
				Value: 2,
			},
			output: &domain.Rule{
				Type:  "minDigit",
				Value: 2,
			},
		},
	}

	for _, t1 := range cases {
		t.Run(t1.input.Rule, func(t *testing.T) {
			res := infra.RuleForObject(t1.input)
			if res.Type != t1.output.Type && res.Value != t1.output.Value {
				t.Fatalf("expect %v, got %v", res, res)
			}
		})
	}

}
