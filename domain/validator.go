package domain

import (
	"regexp"
)

const (
	lowerCaseRegex        = "[a-z]"
	upperCaseRegex        = "[A-Z]"
	digitCharacterRegex   = "[0-9]"
	specialCharacterRegex = "[^[A-Za-zÀ-ȕ\\d\\s]"
)

func MinSize(s string, min int) bool {
	return len(s) >= min
}

func removePerCase(s string, regexp *regexp.Regexp) int {
	replacedString := regexp.FindAllStringSubmatch(s, -1)
	return len(replacedString)
}

func minCase(s, exp string, min int) (bool, error) {
	regex, err := regexp.Compile(exp)
	if err != nil {
		return false, err
	}

	result := removePerCase(s, regex)

	return result >= min, nil
}

func MinUpperCase(s string, min int) (bool, error) {
	return minCase(s, upperCaseRegex, min)
}
func MinLowerCase(s string, min int) (bool, error) {
	return minCase(s, lowerCaseRegex, min)
}
func MinDigit(s string, min int) (bool, error) {
	return minCase(s, digitCharacterRegex, min)
}
func MinSpecialChars(s string, min int) (bool, error) {
	return minCase(s, specialCharacterRegex, min)
}

func NoRepeat(s string) bool {
	data := []byte(s)
	for i := 0; i < len(data); i++ {
		if i > 0 {
			if data[i] == data[i-1] {
				return false
			}
		}
	}
	return true
}
