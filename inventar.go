package inventar

import (
	"context"
)

// UserService is interface of user service
type UserService interface {
	Signup(ctx context.Context, credential *Credential) (bool, error)
	Signin(ctx context.Context, credential *Credential) (bool, error)
}

// UserRepository is interface of user repository
type UserRepository interface {
	Signup(ctx context.Context, credential *Credential) (bool, error)
	Signin(ctx context.Context, credential *Credential) (bool, error)
	GetByUsername(ctx context.Context, username string) (*Credential, error)
}
