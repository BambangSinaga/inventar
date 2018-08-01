package credential

import (
	"context"
	"fmt"
	"time"

	"github.com/West-Labs/inventar"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository inventar.UserRepository
	validator  inventar.Validate
	timeout    time.Duration
}

func (s *service) Signup(ctx context.Context, credential *inventar.Credential) (bool, error) {
	if ctx == nil {
		return false, inventar.ErrContextNil
	}

	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.validator.ValidateStruct(credential); err != nil {
		return false, err
	}

	checkUsername, err := s.repository.GetByUsername(ctx, credential.Username)

	if err != nil {
		return false, err
	}

	if checkUsername.Username != "" {
		return false, inventar.ErrUsernameHasBeenTaken
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credential.Password), 8)

	if err != nil {
		return false, err
	}

	credential.Password = string(hashedPassword)

	res, err := s.repository.Signup(ctx, credential)

	if err != nil {
		return false, err
	}

	return res, nil
}

func (s *service) Signin(ctx context.Context, credential *inventar.Credential) (bool, error) {
	if ctx == nil {
		return false, inventar.ErrContextNil
	}

	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.validator.ValidateStruct(credential); err != nil {
		return false, err
	}

	getByUsername, err := s.repository.GetByUsername(ctx, credential.Username)

	if err != nil {
		return false, inventar.ErrUserNotAuthorized
	}

	if err = bcrypt.CompareHashAndPassword([]byte(getByUsername.Password), []byte(credential.Password)); err != nil {
		fmt.Println(err)
		return false, inventar.ErrUserNotAuthorized
	}

	return true, nil
}

// NewService is implementation of user service interface
func NewService(repository inventar.UserRepository, validator inventar.Validate, timeout time.Duration) inventar.UserService {
	return &service{
		repository: repository,
		validator:  validator,
		timeout:    timeout,
	}
}
