package domain_test

import (
	"psvalidator/domain"
	"testing"
)

func TestLowerCase(t *testing.T) {

	cases := []struct {
		desc, text string
		expected   int
		result     bool
		err        error
	}{
		{desc: "String with 1 lowercase", text: "oNEPASSWORD_123@AA", expected: 1, result: true, err: nil},
		{desc: "String with only lowercase", text: "onepassword_123@aa", expected: 13, result: true, err: nil},
		{desc: "String with 2 lowercases", text: "oNPASSWORD_123@Aa", expected: 2, result: true, err: nil},
		{desc: "String with 3 lowercases but test with 2", text: "oNpASSWORD_123@Aa", expected: 2, result: true, err: nil},
		{desc: "String with 1 lowercase but test 2 two", text: "oNEPASSWORD_123@AA", expected: 2, result: false, err: nil},
		{desc: "String with 0 lowercase", text: "ONEPASSWORD_123@AA", expected: 0, result: true, err: nil},
		{desc: "empty string", text: "", expected: 0, result: true, err: nil},
		{desc: "empty string but test with 2", text: "", expected: 0, result: true, err: nil},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			result, err := domain.MinLowerCase(t1.text, t1.expected)
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
		expected   int
		result     bool
		err        error
	}{
		{desc: "String with 1 uppercase", text: "one_Password_123@aa", expected: 1, result: true, err: nil},
		{desc: "String with only uppercase", text: "ONEPASSWORD_123@AA", expected: 13, result: true, err: nil},
		{desc: "String with 2 uppercase", text: "onPassword_123@aA", expected: 2, result: true, err: nil},
		{desc: "String with 3 uppercase but test with 2", text: "OnPassword_123@aA", expected: 2, result: true, err: nil},
		{desc: "String with 1 uppercase but test 2 two", text: "one_Password_123@aa", expected: 2, result: false, err: nil},
		{desc: "String with 0 uppercase", text: "onepassword_123@aa", expected: 0, result: true, err: nil},
		{desc: "empty string", text: "", expected: 0, result: true, err: nil},
		{desc: "empty string but test with 2", text: "", expected: 0, result: true, err: nil},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			result, err := domain.MinUpperCase(t1.text, t1.expected)
			if err != t1.err {
				t.Fatalf("expect %v , got %v", t1.err, err)
			}
			if result != t1.result {
				t.Fatalf("expect %v , got %v", t1.result, result)
			}
		})
	}
}
