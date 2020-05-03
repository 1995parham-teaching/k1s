package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Status bool
}

func NewHealth() *Health {
	return &Health{Status: true}
}

func (h *Health) Die(c echo.Context) error {
	h.Status = false
	return c.NoContent(http.StatusNoContent)
}

func (h *Health) isAlive(c echo.Context) error {
	if !h.Status {
		time.Sleep(1 * time.Minute)
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Health) Register(g *echo.Group) {
	g.GET("/healthz", h.isAlive)
	g.GET("/die", h.Die)
}
