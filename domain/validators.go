package domain

import (
	"errors"
	"regexp"
)

var (
	LoweError             = errors.New("not valid minimun lowercase rule")
	UpperError            = errors.New("not valid minimun uppercase rule")
	DigitError            = errors.New("not valid minimun digit rule")
	SpecialCharacterError = errors.New("not valid minimun special character rule")
	SizeError             = errors.New("not valid minimun size rule")
	RepeatError           = errors.New("not valid repeat rule")
)

func MinSize(s string, min int) error {
	if len(s) < min {
		return SizeError
	}
	return nil
}

func minCase(s string, rule *Rule) error {
	regex, err := regexp.Compile(RuleRegexp[rule.Type])
	if err != nil {
		return err
	}
	matches := regex.FindAllStringSubmatch(s, -1)

	if len(matches) < rule.Value {
		return RuleErrors[rule.Type]
	}
	return nil
}

func MinUpperCase(s string, min int) error {
	return minCase(s, &Rule{
		Type:  UpperRule,
		Value: min,
	})
}
func MinLowerCase(s string, min int) error {
	return minCase(s, &Rule{
		Type:  LowerRule,
		Value: min,
	})
}
func MinDigit(s string, min int) error {
	return minCase(s, &Rule{
		Type:  DigitRule,
		Value: min,
	})
}
func MinSpecialChars(s string, min int) error {
	return minCase(s, &Rule{
		Type:  SpecialRule,
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