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

// Sign In
func (u *UserRepository) Signin(ctx context.Context, credential *inventar.Credential) (bool, error) {

	query := "SELECT id from user where username = ? and password = ?"
	rows, err := u.DB.Query(query, credential.Username, credential.Password)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	id := 0
	for rows.Next() {

		err = rows.Scan(
			&id,
		)

		if err != nil {
			return false, err
		}
	}

	err = rows.Err()
	if err != nil {
		return false, err
	}

	if id == 0 {
		return false, nil
	}

	return true, nil
}
