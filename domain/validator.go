package domain

import "regexp"

const lowerCaseRegex = "[a-z]"
const upperCaseRegex = "[A-Z]"

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
