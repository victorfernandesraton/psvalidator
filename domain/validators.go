package domain

import (
	"regexp"
)

func MinSize(s string, min int) error {
	if len(s) < min {
		return SizeError
	}
	return nil
}

func minCase(s string, rule *Rule) error {
	regex, err := regexp.Compile(RuleRegexp[rule.Rule])
	if err != nil {
		return err
	}
	matches := regex.FindAllStringSubmatch(s, -1)

	if len(matches) < rule.Value {
		return RuleErrors[rule.Rule]
	}
	return nil
}

func MinUpperCase(s string, min int) error {
	return minCase(s, &Rule{
		Rule:  UpperRule,
		Value: min,
	})
}
func MinLowerCase(s string, min int) error {
	return minCase(s, &Rule{
		Rule:  LowerRule,
		Value: min,
	})
}
func MinDigit(s string, min int) error {
	return minCase(s, &Rule{
		Rule:  DigitRule,
		Value: min,
	})
}
func MinSpecialChars(s string, min int) error {
	return minCase(s, &Rule{
		Rule:  SpecialRule,
		Value: min,
	})
}

func NoRepeat(s string) error {
	data := []byte(s)
	for i := 0; i < len(data); i++ {
		if i > 0 {
			if data[i] == data[i-1] {
				return RepeatError
			}
		}
	}
	return nil
}
