package db

import (
	"database/sql"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	driverSql "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

// MySQLSuite struct for MySQL Suite
type MySQLSuite struct {
	suite.Suite
	ContainerID string
	DSN         string
	DBConn      *sql.DB
	M           *Migration
}

var (
	// MySQLHost holds mysql value
	MySQLHost = `127.0.0.1`
	// MySQLExposedPort holds mysql exposed public port
	MySQLExposedPort = `33060`
	// MySQLUser holds mysql database user
	MySQLUser = `root`
	// MySQLRootPassword holds mysql database password
	MySQLRootPassword = `pass`
	// MySQLDatabase holds mysql database name
	MySQLDatabase = `inventory_db`

	DSN = MySQLUser + `:` + MySQLRootPassword + `@tcp(` + MySQLHost + `:` + MySQLExposedPort + `)/` + MySQLDatabase
)

// SetupSuite setup at the beginning of test
func (s *MySQLSuite) SetupSuite() {
	DisableLogging()

	var err error
	s.DSN = DSN

	s.DBConn, err = sql.Open("mysql", s.DSN+`?parseTime=true`)
	err = s.DBConn.Ping()
	require.NoError(s.T(), err)
	s.DBConn.Exec("set global sql_mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';")
	s.DBConn.Exec("set session sql_mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';")

	s.M, err = RunMigration(s.DSN)
	require.NoError(s.T(), err)
}

// TearDownSuite teardown at the end of test
func (s *MySQLSuite) TearDownSuite() {
	s.DBConn.Close()
}

func DisableLogging() {
	nopLogger := NopLogger{}
	driverSql.SetLogger(nopLogger)
}

type NopLogger struct {
}

func (l NopLogger) Print(v ...interface{}) {
}
