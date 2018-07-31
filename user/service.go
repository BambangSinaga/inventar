package user

import (
	"context"
	"time"

	"github.com/West-Labs/inventar"
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

	res, err := s.repository.Signin(ctx, credential)

	if err != nil {
		return false, err
	}

	return res, nil
}

// NewService is implementation of user service interface
func NewService(repository inventar.UserRepository, validator inventar.Validate, timeout time.Duration) inventar.UserService {
	return &service{
		repository: repository,
		validator:  validator,
		timeout:    timeout,
	}
}
