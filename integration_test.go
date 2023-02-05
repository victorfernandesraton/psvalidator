package main_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/victorfernandesraton/psvalidator/command"
	"github.com/victorfernandesraton/psvalidator/infra"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

	expect := struct {
		err        error
		result     *command.VerifyPasswordCommandResponse
		statusCode int
	}{
		statusCode: http.StatusOK,
		result: &command.VerifyPasswordCommandResponse{
			Verify:  false,
			NoMatch: []string{"minSize"},
		},
	}

	cmd := &command.VerifyPasswordCommand{}
	bodyJSON := `{ "password": "!1234&", "rules": [ { "rule": "minDigit", "value": 4 }, {"rule": "minSize","value": 8} ] }`

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
