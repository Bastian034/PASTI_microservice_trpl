package pindah

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	g.PATCH("", h.Update)
	g.DELETE("/:id", h.Delete)
}
