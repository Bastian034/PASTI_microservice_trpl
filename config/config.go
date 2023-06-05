package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ENV_MODE_PROD = "prod"
	ENV_MODE_DEV  = "dev"
	ENV_MODE_TEST = "test"
)

var Env Configs

type Configs struct {
	Server struct {
		Name string `mapstructure:"APP_NAME" required:"true"`
		Host string `mapstructure:"APP_HOST" required:"true"`
		Port string `mapstructure:"APP_PORT" required:"true"`
		Mode string `mapstructure:"APP_MODE" required:"true"`
	}

	SqlBaptis struct {
		Driver   string `mapstructure:"BAPTIS_SQL_DRIVER" required:"false"`
		Host     string `mapstructure:"BAPTIS_SQL_HOST" required:"false"`
		Port     string `mapstructure:"BAPTIS_SQL_PORT" required:"false"`
		User     string `mapstructure:"BAPTIS_SQL_USER" required:"false"`
		Password string `mapstructure:"BAPTIS_SQL_PASSWORD" required:"false"`
		DbName   string `mapstructure:"BAPTIS_SQL_DBNAME" required:"false"`
	}

	SqlPindah struct {
		Driver   string `mapstructure:"PINDAH_SQL_DRIVER" required:"false"`
		Host     string `mapstructure:"PINDAH_SQL_HOST" required:"false"`
		Port     string `mapstructure:"PINDAH_SQL_PORT" required:"false"`
		User     string `mapstructure:"PINDAH_SQL_USER" required:"false"`
		Password string `mapstructure:"PINDAH_SQL_PASSWORD" required:"false"`
		DbName   string `mapstructure:"PINDAH_SQL_DBNAME" required:"false"`
	}

	SqlPemberkatan struct {
		Driver   string `mapstructure:"PEMBERKATAN_SQL_DRIVER" required:"false"`
		Host     string `mapstructure:"PEMBERKATAN_SQL_HOST" required:"false"`
		Port     string `mapstructure:"PEMBERKATAN_SQL_PORT" required:"false"`
		User     string `mapstructure:"PEMBERKATAN_SQL_USER" required:"false"`
		Password string `mapstructure:"PEMBERKATAN_SQL_PASSWORD" required:"false"`
		DbName   string `mapstructure:"PEMBERKATAN_SQL_DBNAME" required:"false"`
	}
}

func LoadConfig(appMode string) *Configs {
	if appMode == ENV_MODE_DEV {
		return loadENV(".env.dev")
	} else if appMode == ENV_MODE_PROD {
		return loadENV(".env.prod")

	} else if appMode == ENV_MODE_TEST {
		return loadENV(".env.test")
	}
	return loadENV(".env.dev")
}

func loadENV(envName string) *Configs {
	viper.AddConfigPath(".")
	viper.SetConfigName(envName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
		})
	}

	err = viper.Unmarshal(&Env.SqlBaptis)
	err = viper.Unmarshal(&Env.SqlPemberkatan)
	err = viper.Unmarshal(&Env.SqlPindah)
	err = viper.Unmarshal(&Env.Server)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
		})
	}

	return &Env
}
