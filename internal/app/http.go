package app

import (
	"fmt"
	"gereja-services/config"
	"gereja-services/docs"
	"gereja-services/internal/middleware"
	"gereja-services/internal/repository/interfaces"
	"gereja-services/internal/repository/mysql"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg             *config.Configs
	log             *logrus.Logger
	echo            *echo.Echo
	repoBaptis      interfaces.BaptisRepository
	repoPindah      interfaces.PindahRepository
	repoPemberkatan interfaces.PemberkatanRepository
}

func New(appMode string) *App {
	cfg := config.LoadConfig(appMode)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	log := logrus.New()

	echo := echo.New()

	middleware.Init(echo)

	repo := &App{
		repoBaptis:      mysql.NewRepoBaptis(log, cfg),
		repoPindah:      mysql.NewRepoPindah(log, cfg),
		repoPemberkatan: mysql.NewRepoPemberkatan(log, cfg),
	}

	Router(echo, repo, log)

	logrus.Fatal(echo.Start(":" + cfg.Server.Port))

	return &App{
		cfg:  cfg,
		log:  log,
		echo: echo,
	}
}
