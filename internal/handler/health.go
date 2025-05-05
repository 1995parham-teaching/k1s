package handler

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Health struct {
	Status bool
}

func NewHealth() *Health {
	return &Health{Status: true}
}

// nolint: wrapcheck
func (h *Health) Die(c *fiber.Ctx) error {
	h.Status = false

	return c.Status(http.StatusNoContent).Send(nil)
}

// nolint: wrapcheck
func (h *Health) IsAlive(c *fiber.Ctx) error {
	if !h.Status {
		time.Sleep(1 * time.Minute)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusNoContent).Send(nil)
}

func (h *Health) Register(g fiber.Router) {
	g.Get("/healthz", h.IsAlive)
	g.Get("/die", h.Die)
}
