package domain_test

import (
	"psvalidator/domain"
	"testing"
)

func TestSize(t *testing.T) {
	cases := []struct {
		desc, text string
		minimun    int
		result     bool
	}{
		{desc: "string with 3 characters", text: "a@1", minimun: 3, result: true},
		{desc: "string with 3 characters but test for 2", text: "a@1", minimun: 2, result: true},
		{desc: "string with 3 characters but test for 4", text: "a@1", minimun: 4, result: false},
		{desc: "empty string", text: "", minimun: 0, result: true},
		{desc: "empty string but test for 4", text: "", minimun: 4, result: false},
	}
	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			result := domain.MinSize(t1.text, t1.minimun)
			if result != t1.result {
				t.Fatalf("expect %v , got %v", t1.result, result)
			}
		})
	}
}

func TestLowerCase(t *testing.T) {

	cases := []struct {
		desc, text string
		minimun    int
		result     bool
		err        error
	}{
		{desc: "String with 1 lowercase", text: "oNEPASSWORD_123@AA", minimun: 1, result: true, err: nil},
		{desc: "String with only lowercase", text: "onepassword_123@aa", minimun: 13, result: true, err: nil},
		{desc: "String with 2 lowercases", text: "oNPASSWORD_123@Aa", minimun: 2, result: true, err: nil},
		{desc: "String with 3 lowercases but test with 2", text: "oNpASSWORD_123@Aa", minimun: 2, result: true, err: nil},
		{desc: "String with 1 lowercase but test 2", text: "oNEPASSWORD_123@AA", minimun: 2, result: false, err: nil},
		{desc: "String with 0 lowercase", text: "ONEPASSWORD_123@AA", minimun: 0, result: true, err: nil},
		{desc: "empty string", text: "", minimun: 0, result: true, err: nil},
		{desc: "empty string but test with 2", text: "", minimun: 2, result: false, err: nil},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			result, err := domain.MinLowerCase(t1.text, t1.minimun)
			if err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
			if result != t1.result {
				t.Fatalf("expect %v , got %v", t1.result, result)
			}
		})
	}
}

func TestUpperCase(t *testing.T) {

	cases := []struct {
		desc, text string
		minimun    int
		result     bool
		err        error
	}{
		{desc: "String with 1 uppercase", text: "one_Password_123@aa", minimun: 1, result: true, err: nil},
		{desc: "String with only uppercase", text: "ONEPASSWORD_123@AA", minimun: 13, result: true, err: nil},
		{desc: "String with 2 uppercase", text: "onPassword_123@aA", minimun: 2, result: true, err: nil},
		{desc: "String with 3 uppercase but test with 2", text: "OnPassword_123@aA", minimun: 2, result: true, err: nil},
		{desc: "String with 1 uppercase but test 2", text: "one_Password_123@aa", minimun: 2, result: false, err: nil},
		{desc: "String with 0 uppercase", text: "onepassword_123@aa", minimun: 0, result: true, err: nil},
		{desc: "empty string", text: "", minimun: 0, result: true, err: nil},
		{desc: "empty string but test with 2", text: "", minimun: 2, result: false, err: nil},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			result, err := domain.MinUpperCase(t1.text, t1.minimun)
			if err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
			if result != t1.result {
				t.Fatalf("expect %v , got %v", t1.result, result)
			}
		})
	}
}

func TestDigit(t *testing.T) {

	cases := []struct {
		desc, text string
		minimun    int
		result     bool
		err        error
	}{
		{desc: "String with 1 digit", text: "one_Password_1@aa", minimun: 1, result: true, err: nil},
		{desc: "String with only digit", text: "123", minimun: 3, result: true, err: nil},
		{desc: "String with 2 digit", text: "onPassword_12@aA", minimun: 2, result: true, err: nil},
		{desc: "String with 3 digit but test with 2", text: "OnPassword_123@aA", minimun: 2, result: true, err: nil},
		{desc: "String with 1 digit but test 2", text: "one_Password_1@aa", minimun: 2, result: false, err: nil},
		{desc: "String with 0 digit", text: "onepassword_@aa", minimun: 0, result: true, err: nil},
		{desc: "empty string", text: "", minimun: 0, result: true, err: nil},
		{desc: "empty string but test with 2", text: "", minimun: 2, result: false, err: nil},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			result, err := domain.MinDigit(t1.text, t1.minimun)
			if err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
			if result != t1.result {
				t.Fatalf("expect %v , got %v", t1.result, result)
			}
		})
	}
}

func TestSpecialChars(t *testing.T) {

	cases := []struct {
		desc, text string
		minimun    int
		result     bool
		err        error
	}{
		{desc: "String with 1 special character", text: "onePassword1@aa", minimun: 1, result: true, err: nil},
		{desc: "String with only special character", text: "-_-", minimun: 3, result: true, err: nil},
		{desc: "String with 2 special character", text: "onPassword_12@aA", minimun: 2, result: true, err: nil},
		{desc: "String with 3 special character but test with 2", text: "On+Password_123@aA", minimun: 2, result: true, err: nil},
		{desc: "String with 1 special character but test 2", text: "onePassword1@aa", minimun: 2, result: false, err: nil},
		{desc: "String with 0 special character", text: "OnePassword123aA", minimun: 0, result: true, err: nil},
		{desc: "empty string", text: "", minimun: 0, result: true, err: nil},
		{desc: "empty string but test with 2", text: "", minimun: 2, result: false, err: nil},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			result, err := domain.MinSpecialChars(t1.text, t1.minimun)
			if err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
			if result != t1.result {
				t.Fatalf("expect %v , got %v", t1.result, result)
			}
		})
	}
}

func TestNoRepeat(t *testing.T) {

	cases := []struct {
		desc, text string
		result     bool
		err        error
	}{
		{desc: "String with not sequential repetition", text: "aba123", result: true, err: nil},
		{desc: "String with with 1 sequential repetition", text: "aab123", result: false, err: nil},
		{desc: "String with with 3 sequential repetition", text: "abaabcbc1123", result: false, err: nil},
		{desc: "empty string", text: "", result: true, err: nil},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			result := domain.NoRepeat(t1.text)
			if result != t1.result {
				t.Fatalf("expect %v , got %v", t1.result, result)
			}
		})
	}
}
