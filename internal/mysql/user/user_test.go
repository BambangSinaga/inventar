package user

import (
	"context"
	"testing"
	"time"

	"github.com/West-Labs/inventar"

	"github.com/West-Labs/inventar/internal/mysql/db"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	db.MySQLSuite
}

func TestUserMysqlRepositoryTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip(`Skip user repository test`)
	}

	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) SetupTest() {
	logrus.Info("Starting to Migrate Up Data")
	errs, ok := s.M.Up()
	require.True(s.T(), ok)
	require.Len(s.T(), errs, 0)
}

func (s *UserRepositoryTestSuite) TearDownTest() {
	logrus.Info("Starting to Migrate Down Data")
	errs, ok := s.M.Down()
	require.True(s.T(), ok)
	require.Len(s.T(), errs, 0)
}

func (s *UserRepositoryTestSuite) seedUser(c *inventar.Credential) error {
	query := "INSERT INTO `user` (`username`, `password`, `create_time`) VALUES (?, ?, ?)"
	preparedQuery, err := s.DBConn.Prepare(query)
	require.NoError(s.T(), err)
	defer preparedQuery.Close()

	res, insertError := preparedQuery.Exec(c.Username, c.Password, time.Now())
	require.NoError(s.T(), insertError)

	_, err = res.LastInsertId()
	require.NoError(s.T(), err)

	return nil
}

func (s *UserRepositoryTestSuite) TestSigin() {

	mockUser := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	err := s.seedUser(mockUser)
	require.NoError(s.T(), err)

	repo := UserMysqlRepository{DB: s.DBConn}
	ctx := context.Background()
	res, err := repo.Signin(ctx, mockUser)
	assert.NoError(s.T(), err)
	assert.True(s.T(), res)
}

func (s *UserRepositoryTestSuite) TestSiginInvalidCredential() {

	mockUser := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	mockLogin := &inventar.Credential{
		Username: "admin",
		Password: "admin",
	}

	err := s.seedUser(mockUser)
	require.NoError(s.T(), err)

	repo := UserMysqlRepository{DB: s.DBConn}
	ctx := context.Background()
	res, err := repo.Signin(ctx, mockLogin)
	assert.NoError(s.T(), err)
	assert.False(s.T(), res)
}

func (s *UserRepositoryTestSuite) TestSignup() {

	mockUser := &inventar.Credential{
		Username: "admin",
		Password: "admin123",
	}

	repo := UserMysqlRepository{DB: s.DBConn}
	ctx := context.Background()
	res, err := repo.Signup(ctx, mockUser)
	assert.NoError(s.T(), err)
	assert.True(s.T(), res)
}
