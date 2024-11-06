package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"mahestan/conf"
	"mahestan/internal/controller/dispatcher"
	"mahestan/internal/repository/sqlite"
)

func runServer(router *gin.Engine, cfg *conf.AppConfig) {
	port := fmt.Sprintf(":%d", cfg.App.Port)
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				conf.GetMahestanRuntimeConfig,
				sqlite.ConnectToSqlite,
				dispatcher.SetupRouter,
			),
			fx.Invoke(runServer),
		),
	).Run()
}
