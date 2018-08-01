package main

import (
	"os"

	"github.com/West-Labs/inventar/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.RootCMD.Execute(); err != nil {
		logrus.Errorln(err.Error())
		os.Exit(1)
	}
}
