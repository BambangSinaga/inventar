package credential

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/West-Labs/inventar"
	"github.com/West-Labs/inventar/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var (
	timeout          = time.Duration(5)
	subscriptionType = "type"
)

func TestSignup(t *testing.T) {

	mockCredential := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	r := new(mocks.UserRepository)
	r.On("GetByUsername", mock.Anything, mockCredential.Username).Return(&inventar.Credential{}, nil)
	r.On("Signup", mock.Anything, mockCredential).Return(true, nil)

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signup(context.TODO(), mockCredential)

	assert.NoError(t, err)
	assert.True(t, res)
}

func TestSignupWithContextNil(t *testing.T) {

	mockCredential := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	r := new(mocks.UserRepository)

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signup(nil, mockCredential)

	assert.Error(t, err)
	assert.False(t, res)
}

func TestSignupWithErrorValidate(t *testing.T) {

	mockCredential := &inventar.Credential{
		Username: "admin",
	}

	r := new(mocks.UserRepository)

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signup(context.TODO(), mockCredential)

	assert.Error(t, err)
	assert.False(t, res)
}

func TestSignupWithError(t *testing.T) {

	mockCredential := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	r := new(mocks.UserRepository)
	r.On("GetByUsername", mock.Anything, mockCredential.Username).Return(&inventar.Credential{}, nil)
	r.On("Signup", mock.Anything, mockCredential).Return(false, errors.New("Error"))

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signup(context.TODO(), mockCredential)

	assert.Error(t, err)
	assert.False(t, res)
}

func TestSignupErrorWithUsernamehasbeentaken(t *testing.T) {

	mockCredential := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	r := new(mocks.UserRepository)
	r.On("GetByUsername", mock.Anything, mockCredential.Username).Return(mockCredential, nil)
	r.On("Signup", mock.Anything, mockCredential).Return(false, nil)

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signup(context.TODO(), mockCredential)
	assert.Error(t, err)
	assert.False(t, res)
}

func TestSignin(t *testing.T) {

	mockCredentialNonHashed := &inventar.Credential{
		Username: "admin",
		Password: "tesla123@",
	}

	mockCredentialHashed := &inventar.Credential{
		Username: "admin",
		Password: "tesla123@",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mockCredentialHashed.Password), 8)

	mockCredentialHashed.Password = string(hashedPassword)
	fmt.Println(mockCredentialHashed.Password)

	r := new(mocks.UserRepository)
	r.On("GetByUsername", mock.Anything, mockCredentialHashed.Username).Return(mockCredentialHashed, nil)

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signin(context.TODO(), mockCredentialNonHashed)
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestSigninWitContextNil(t *testing.T) {

	mockCredential := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	r := new(mocks.UserRepository)

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signin(nil, mockCredential)

	assert.Error(t, err)
	assert.False(t, res)
}

func TestSigninWithErrorValidate(t *testing.T) {

	mockCredential := &inventar.Credential{
		Username: "admin",
	}

	r := new(mocks.UserRepository)

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signin(context.TODO(), mockCredential)

	assert.Error(t, err)
	assert.False(t, res)
}

func TestSigninWithInvalidUsernamePassoword(t *testing.T) {

	mockCredential := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	r := new(mocks.UserRepository)
	r.On("GetByUsername", mock.Anything, mockCredential.Username).Return(&inventar.Credential{}, nil)

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signin(context.TODO(), mockCredential)

	assert.Error(t, err)
	assert.False(t, res)
}

func TestSigninWithNoRecord(t *testing.T) {

	mockCredentialNonHashed := &inventar.Credential{
		Username: "admin",
		Password: "tesla123@",
	}

	mockCredentialHashed := &inventar.Credential{
		Username: "admin",
		Password: "tesla123@",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mockCredentialHashed.Password), 8)

	mockCredentialHashed.Password = string(hashedPassword)
	fmt.Println(mockCredentialHashed.Password)

	r := new(mocks.UserRepository)
	r.On("GetByUsername", mock.Anything, mockCredentialHashed.Username).Return(nil, errors.New("Error"))

	v := inventar.NewValidator()

	s := NewService(r, v, timeout)

	res, err := s.Signin(context.TODO(), mockCredentialNonHashed)
	assert.Error(t, err)
	assert.False(t, res)
}
