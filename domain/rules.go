package domain

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
	Type  RuleEnum `json:"rule"`
	Value int      `json:"value"`
}
