package credential

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/West-Labs/inventar"

	"github.com/West-Labs/inventar/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignup(t *testing.T) {
	var mockRequestCredential inventar.Credential

	jsonMockCredential := `
		{
			"username": "admin",
			"password": "admin123"
		}`

	require.NoError(t, json.Unmarshal([]byte(jsonMockCredential), &mockRequestCredential))

	u := new(mocks.UserService)
	u.On("Signup", mock.Anything, &mockRequestCredential).Return(true, nil)

	h := httpHandler{u}
	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/user/signup", strings.NewReader(jsonMockCredential))
	require.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	h.HandleSignupUser(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)

	u.AssertExpectations(t)
}

func TestSignupWithInternalServerError(t *testing.T) {
	var mockRequestCredential inventar.Credential

	jsonMockCredential := `
		{
			"username": "admin",
			"password": "admin123"
		}`

	require.NoError(t, json.Unmarshal([]byte(jsonMockCredential), &mockRequestCredential))

	u := new(mocks.UserService)
	u.On("Signup", mock.Anything, &mockRequestCredential).Return(false, errors.New("Error"))

	h := httpHandler{u}
	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/user/signup", strings.NewReader(jsonMockCredential))
	require.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	h.HandleSignupUser(ctx)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	u.AssertExpectations(t)
}

func TestSignupWithInvalidData(t *testing.T) {
	jsonMockCredential := `
		{
			"username": "admin",
		}`

	u := new(mocks.UserService)

	h := httpHandler{u}
	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/user/signup", strings.NewReader(jsonMockCredential))
	require.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	h.HandleSignupUser(ctx)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestSignin(t *testing.T) {
	var mockRequestCredential inventar.Credential

	jsonMockCredential := `
		{
			"username": "admin",
			"password": "admin123"
		}`

	require.NoError(t, json.Unmarshal([]byte(jsonMockCredential), &mockRequestCredential))

	u := new(mocks.UserService)
	u.On("Signin", mock.Anything, &mockRequestCredential).Return(true, nil)

	h := httpHandler{u}
	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/user/signin", strings.NewReader(jsonMockCredential))
	require.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	h.HandleSigninUser(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)

	u.AssertExpectations(t)
}

func TestSigninWithInternalServerError(t *testing.T) {
	var mockRequestCredential inventar.Credential

	jsonMockCredential := `
		{
			"username": "admin",
			"password": "admin123"
		}`

	require.NoError(t, json.Unmarshal([]byte(jsonMockCredential), &mockRequestCredential))

	u := new(mocks.UserService)
	u.On("Signin", mock.Anything, &mockRequestCredential).Return(false, errors.New("Error"))

	h := httpHandler{u}
	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/user/signin", strings.NewReader(jsonMockCredential))
	require.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	h.HandleSigninUser(ctx)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	u.AssertExpectations(t)
}

func TestSigninWithInvalidData(t *testing.T) {
	jsonMockCredential := `
		{
			"username": "admin",
		}`

	u := new(mocks.UserService)

	h := httpHandler{u}
	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/user/signin", strings.NewReader(jsonMockCredential))
	require.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	h.HandleSigninUser(ctx)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}
