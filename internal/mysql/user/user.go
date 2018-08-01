package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/West-Labs/inventar"
)

// UserMysqlRepository is struct to initiate mysql repo for user
type UserRepository struct {
	DB *sql.DB
}

// Sign Up
func (u *UserRepository) Signup(ctx context.Context, credential *inventar.Credential) (bool, error) {

	trx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}

	// Insert Into User
	query := "INSERT INTO `user` " +
		"(`username`, `password`, `create_time`) " +
		"VALUES (?, ?, ?)"

	stmt, err := trx.PrepareContext(ctx, query)
	if err != nil {
		trx.Rollback()
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, credential.Username, credential.Password, time.Now())
	if err != nil {
		trx.Rollback()
		return false, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		trx.Rollback()
		return false, err
	}

	return true, trx.Commit()
}

// Get By Username
func (u *UserRepository) GetByUsername(ctx context.Context, username string) (*inventar.Credential, error) {

	query := "SELECT username, password from user where username = ?"
	rows, err := u.DB.Query(query, username)
	if err != nil {
		return &inventar.Credential{}, err
	}
	defer rows.Close()

	cred := &inventar.Credential{}
	for rows.Next() {

		err = rows.Scan(
			&cred.Username,
			&cred.Password,
		)

		if err != nil {
			return &inventar.Credential{}, err
		}
	}

	err = rows.Err()
	if err != nil {
		return &inventar.Credential{}, err
	}

	if cred.Username == "" {
		return &inventar.Credential{}, err
	}

	return cred, err
}
