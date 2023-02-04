package domain_test

import (
	"psvalidator/domain"
	"testing"
)

func TestSize(t *testing.T) {
	cases := []struct {
		desc, text string
		minimun    int
		err        error
	}{
		{desc: "string with 3 characters", text: "a@1", minimun: 3},
		{desc: "string with 3 characters but test for 2", text: "a@1", minimun: 2},
		{desc: "string with 3 characters but test for 4", text: "a@1", minimun: 4, err: domain.SizeError},
		{desc: "empty string", text: "", minimun: 0},
		{desc: "empty string but test for 4", text: "", minimun: 4, err: domain.SizeError},
	}
	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			if err := domain.MinSize(t1.text, t1.minimun); err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
		})
	}
}

func TestLowerCase(t *testing.T) {

	cases := []struct {
		desc, text string
		minimun    int
		err        error
	}{
		{desc: "String with 1 lowercase", text: "oNEPASSWORD_123@AA", minimun: 1},
		{desc: "String with only lowercase", text: "onepassword_123@aa", minimun: 13},
		{desc: "String with 2 lowercases", text: "oNPASSWORD_123@Aa", minimun: 2},
		{desc: "String with 3 lowercases but test with 2", text: "oNpASSWORD_123@Aa", minimun: 2},
		{desc: "String with 1 lowercase but test 2", text: "oNEPASSWORD_123@AA", minimun: 2, err: domain.LoweError},
		{desc: "String with 0 lowercase", text: "ONEPASSWORD_123@AA", minimun: 0},
		{desc: "empty string", text: "", minimun: 0},
		{desc: "empty string but test with 2", text: "", minimun: 2, err: domain.LoweError},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			if err := domain.MinLowerCase(t1.text, t1.minimun); err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
		})
	}
}

func TestUpperCase(t *testing.T) {

	cases := []struct {
		desc, text string
		minimun    int
		err        error
	}{
		{desc: "String with 1 uppercase", text: "one_Password_123@aa", minimun: 1},
		{desc: "String with only uppercase", text: "ONEPASSWORD_123@AA", minimun: 13},
		{desc: "String with 2 uppercase", text: "onPassword_123@aA", minimun: 2},
		{desc: "String with 3 uppercase but test with 2", text: "OnPassword_123@aA", minimun: 2},
		{desc: "String with 1 uppercase but test 2", text: "one_Password_123@aa", minimun: 2, err: domain.UpperError},
		{desc: "String with 0 uppercase", text: "onepassword_123@aa", minimun: 0},
		{desc: "empty string", text: "", minimun: 0},
		{desc: "empty string but test with 2", text: "", minimun: 2, err: domain.UpperError},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			if err := domain.MinUpperCase(t1.text, t1.minimun); err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
		})
	}
}

func TestDigit(t *testing.T) {

	cases := []struct {
		desc, text string
		minimun    int
		err        error
	}{
		{desc: "String with 1 digit", text: "one_Password_1@aa", minimun: 1},
		{desc: "String with only digit", text: "123", minimun: 3},
		{desc: "String with 2 digit", text: "onPassword_12@aA", minimun: 2},
		{desc: "String with 3 digit but test with 2", text: "OnPassword_123@aA", minimun: 2},
		{desc: "String with 1 digit but test 2", text: "one_Password_1@aa", minimun: 2, err: domain.DigitError},
		{desc: "String with 0 digit", text: "onepassword_@aa", minimun: 0},
		{desc: "empty string", text: "", minimun: 0},
		{desc: "empty string but test with 2", text: "", minimun: 2, err: domain.DigitError},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			if err := domain.MinDigit(t1.text, t1.minimun); err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
		})
	}
}

func TestSpecialChars(t *testing.T) {

	cases := []struct {
		desc, text string
		minimun    int
		err        error
	}{
		{desc: "String with 1 special character", text: "onePassword1@aa", minimun: 1},
		{desc: "String with only special character", text: "-_-", minimun: 3},
		{desc: "String with 2 special character", text: "onPassword_12@aA", minimun: 2},
		{desc: "String with 3 special character but test with 2", text: "On+Password_123@aA", minimun: 2},
		{desc: "String with 1 special character but test 2", text: "onePassword1@aa", minimun: 2, err: domain.SpecialCharacterError},
		{desc: "String with 0 special character", text: "OnePassword123aA", minimun: 0},
		{desc: "empty string", text: "", minimun: 0},
		{desc: "empty string but test with 2", text: "", minimun: 2, err: domain.SpecialCharacterError},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			if err := domain.MinSpecialChars(t1.text, t1.minimun); err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
		})
	}
}

func TestNoRepeat(t *testing.T) {

	cases := []struct {
		desc, text string
		err        error
	}{
		{desc: "String with not sequential repetition", text: "aba123"},
		{desc: "String with with 1 sequential repetition", text: "aab123", err: domain.RepeatError},
		{desc: "String with with 3 sequential repetition", text: "abaabcbc1123", err: domain.RepeatError},
		{desc: "empty string", text: ""},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			if err := domain.NoRepeat(t1.text); err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
		})
	}
}
