package conf

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/sqlite"
)

type AppConfig struct {
	Database struct {
		Sqlite struct {
			Name          string `envconfig:"NAME"`
			MigrationPath string `envconfig:"MIGRATION_PATH"`
		} `envconfig:"SQLITE"`
	} `envconfig:"DATABASE"`

	App struct {
		Debug bool   `envconfig:"DEBUG"`
		Name  string `envconfig:"NAME"`
		Port  int    `envconfig:"PORT"`
	} `envconfig:"APP"`
}

func GetMahestanRuntimeConfig() (*AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.WithError(err).Errorln("dont load .env file!")
	}

	cfg := &AppConfig{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
