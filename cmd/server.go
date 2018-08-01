package cmd

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/West-Labs/inventar"
	_credentialService "github.com/West-Labs/inventar/credential"
	_userDelivery "github.com/West-Labs/inventar/internal/http/credential"
	_userRepository "github.com/West-Labs/inventar/internal/mysql/user"
)

var serverCMD = &cobra.Command{
	Use:   "http",
	Short: "Start http server of europa",
	Run: func(cmd *cobra.Command, args []string) {
		mysqlDB, err := sql.Open("mysql", initConnection())
		if err != nil {
			logrus.Errorln("Failed to connect to database: " + err.Error())
			os.Exit(1)
		}

		e := echo.New()

		e.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})

		initService(mysqlDB, e)

		address := viper.GetString("server.address")
		logrus.Infof("Start Listening on: %v", address)
		e.Start(address)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCMD.AddCommand(serverCMD)
	serverCMD.PersistentFlags().String("config", "", "Set this flag to use a configuration file")
}

func initConfig() {
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.BindPFlag("config", serverCMD.Flags().Lookup("config"))

	if config := viper.GetString("config"); config != "" {
		viper.SetConfigType("json")

		viper.SetConfigFile(config)
		if err := viper.ReadInConfig(); err != nil {
			logrus.Errorln(err.Error())
		}
	}
}

func initService(db *sql.DB, e *echo.Echo) {
	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second
	validator := inventar.NewValidator()
	userRepository := _userRepository.UserRepository{DB: db}
	credentialService := _credentialService.NewService(&userRepository, validator, timeout)
	_userDelivery.Init(e, credentialService)
}

func initConnection() string {
	mysqlUser := viper.GetString("mysql.user")
	mysqlPassword := viper.GetString("mysql.pass")
	mysqlHost := viper.GetString("mysql.host")
	mysqlPort := viper.GetString("mysql.port")
	mysqlDatabaseName := viper.GetString("mysql.name")
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabaseName)
}
