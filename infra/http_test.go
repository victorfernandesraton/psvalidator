package infra_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/victorfernandesraton/psvalidator/command"
	"github.com/victorfernandesraton/psvalidator/infra"
)

func TestBadRequestError(t *testing.T) {

	expect := struct {
		err        error
		result     string
		statusCode int
	}{
		err:        echo.NewHTTPError(http.StatusBadRequest, command.NotValidRuleError),
		statusCode: http.StatusBadRequest,
	}
	bodyJSON := `{ "password": "TesteSenhaForte!123&", "rules": [ { "rule": "otherRuçe", "value": 8 }, { "rule": "minSpecialChars", "value": 2 }, { "rule": "noRepeted", "value": 0 }, { "rule": "minDigit", "value": 4 } ] }`

	cmd := &command.VerifyPasswordCommand{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &infra.VerifyHttpController{
		VerifyCommand: cmd,
	}

	err := h.Handler(c)

	if err.Error() != expect.err.Error() {
		t.Fatalf("expect %v, got %v", expect.err, err)
	}
}

func TestWithSucessValidPass(t *testing.T) {

	expect := struct {
		err        error
		result     *command.VerifyPasswordCommandResponse
		statusCode int
	}{
		err:        nil,
		statusCode: http.StatusOK,
		result: &command.VerifyPasswordCommandResponse{
			Verify:  true,
			NoMatch: []string{},
		},
	}

	cmd := &command.VerifyPasswordCommand{}
	bodyJSON :=
		`{
			"password": "TesteSenhaForte!1234&",
			"rules": [
				{"rule": "minSize","value": 8},
				{"rule": "minSpecialChars","value": 2},
				{"rule": "noRepeat","value": 0},
				{"rule": "minDigit","value": 4}
			]
		}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &infra.VerifyHttpController{
		VerifyCommand: cmd,
	}

	err := h.Handler(c)

	if err != expect.err {
		t.Fatalf("expect %v, got %v", expect.err, err)
	}
	response := new(command.VerifyPasswordCommandResponse)
	if err := json.Unmarshal(rec.Body.Bytes(), response); err != nil {
		t.Fatalf("malformed body response")
	}

	if rec.Code != expect.statusCode {
		t.Fatalf("expect %v, got %v", expect.statusCode, rec.Code)
	}

	if response.Verify != expect.result.Verify {
		t.Fatalf("expect %v, got %v", expect.result, response)
	}
}

func TestWithErrorValidPass(t *testing.T) {

	cmd := &command.VerifyPasswordCommand{}
	e := echo.New()

	cases := []struct {
		body    string
		desc    string
		moMatch []string
	}{
		{
			desc: "invalid min size",
			body: `{ 
				"password": "!1234&", 
				"rules": [ 
					{ "rule": "minDigit", "value": 4 },
					{"rule": "minSize","value": 8} 
				] 
			}`,
			moMatch: []string{"minSize"},
		},
		{
			desc: "invalid min size and mindigit",
			body: `{ 
				"password": "!1234&", 
				"rules": [ 
					{ "rule": "minDigit", "value": 5 },
					{"rule": "minSize","value": 8} 
				] 
			}`,
			moMatch: []string{"minSize", "minDigit"},
		},
	}

	for _, t1 := range cases {
		t.Run(t1.desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(t1.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := &infra.VerifyHttpController{
				VerifyCommand: cmd,
			}

			err := h.Handler(c)

			if err != nil {
				t.Fatalf("expect %v, got %v", nil, err)
			}
			response := new(command.VerifyPasswordCommandResponse)
			if err := json.Unmarshal(rec.Body.Bytes(), response); err != nil {
				t.Fatalf("malformed body response")
			}
			if rec.Code != http.StatusOK {
				t.Fatalf("expect %v, got %v", http.StatusOK, rec.Code)
			}

			if response.Verify {
				t.Fatalf("expect %v, got %v", false, response)
			}
			if len(response.NoMatch) != len(t1.moMatch) {
				t.Fatalf("different size list expect %v, got %v", t1.moMatch, response.NoMatch)
			}
			for _, match := range response.NoMatch {
				findIndex := sort.Search(len(t1.moMatch), func(i int) bool {
					return t1.moMatch[i] == match
				})
				if findIndex > len(t1.moMatch) {
					t.Fatalf("not found %v in %v", match, t1.moMatch)
				}
			}
		})
	}
}
