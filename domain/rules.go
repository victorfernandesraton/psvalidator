package domain

import "errors"

type RuleEnum = string

const (
	SizeRule    RuleEnum = "minSize"
	SpecialRule          = "minSpecialChars"
	UpperRule            = "minUpperCase"
	LowerRule            = "minLowerCase"
	DigitRule            = "minDigit"
	RepeatRule           = "noRepeat"
)

const (
	lowerCaseRegex        = "[a-z]"
	upperCaseRegex        = "[A-Z]"
	digitCharacterRegex   = "[0-9]"
	specialCharacterRegex = "[^[A-Za-zÀ-ȕ\\d\\s]"
)

var (
	LoweError             = errors.New("not valid minimun lowercase rule")
	UpperError            = errors.New("not valid minimun uppercase rule")
	DigitError            = errors.New("not valid minimun digit rule")
	SpecialCharacterError = errors.New("not valid minimun special character rule")
	SizeError             = errors.New("not valid minimun size rule")
	RepeatError           = errors.New("not valid repeat rule")
)

var RuleErrors = map[RuleEnum]error{
	SizeRule:    SizeError,
	UpperRule:   UpperError,
	LowerRule:   LoweError,
	DigitRule:   DigitError,
	SpecialRule: SpecialCharacterError,
	RepeatRule:  RepeatError,
}

var RuleRegexp = map[RuleEnum]string{
	UpperRule:   upperCaseRegex,
	LowerRule:   lowerCaseRegex,
	DigitRule:   digitCharacterRegex,
	SpecialRule: specialCharacterRegex,
}

type Rule struct {
	Rule  RuleEnum `json:"rule"`
	Value int      `json:"value"`
}
