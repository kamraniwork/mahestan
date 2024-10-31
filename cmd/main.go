package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"mahestan/conf"
	"mahestan/internal/controller/dispatcher"
	"mahestan/internal/repository/sqlite"
)

func main() {
	if err := runServer(); err != nil {
		logrus.WithError(err).Errorln("can not run server!")
		return
	}
}

func runServer() error {
	cfg, err := conf.GetMahestanRuntimeConfig()
	if err != nil {
		return err
	}

	db, err := sqlite.ConnectToSqlite(cfg)
	if err != nil {
		return err
	}

	routerEngine, err := dispatcher.SetupRouter(db)
	if err != nil {
		return err
	}

	return routerEngine.Run(fmt.Sprintf(":%d", cfg.App.Port))
}
