package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/titasp/gin-starter-project/config"
	"github.com/titasp/gin-starter-project/database"
	"github.com/titasp/gin-starter-project/server"
)

func main() {

	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		logrus.Fatal("Usage: server -e {mode}")
	}
	flag.Parse()

	err := config.Init(*environment)
	if err != nil {
		logrus.WithField("error", err).Fatal("error on parsing configuration file")
	}

	err = database.Init()
	if err != nil {
		logrus.WithField("error", err).Fatal("failed to initialize database")
	}

	s := server.New()
	logrus.Fatal(s.Run(":" + config.GetConfig().GetString("server.port")))
}
