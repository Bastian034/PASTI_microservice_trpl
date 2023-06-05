package app

import (
	"gereja-services/internal/usecase/baptis"
	"gereja-services/internal/usecase/pemberkatan"
	"gereja-services/internal/usecase/pindah"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func Router(e *echo.Echo, r *App, l *logrus.Logger) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		message := "Welcome to Gereja Services"
		return c.String(http.StatusOK, message)
	})
	baptis.NewHandler(r.repoBaptis, l).Route(e.Group("/baptis"))
	pemberkatan.NewHandler(r.repoPemberkatan, l).Route(e.Group("/pemberkatan"))
	pindah.NewHandler(r.repoPindah, l).Route(e.Group("/pindah"))
}
