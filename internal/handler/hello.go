package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Hello handles the hello messages.
type Hello struct {
	Hostname        string
	GreetingMessage string

	Logger *zap.Logger
}

// NewHello creates a new instance of hello handler.
func NewHello(msg string, logger *zap.Logger) *Hello {
	hostname, err := os.Hostname()
	if err != nil {
		logger.Error("cannot detect host name", zap.Error(err))

		hostname = "parham"
	}

	return &Hello{
		Hostname:        hostname,
		GreetingMessage: msg,

		Logger: logger,
	}
}

// Say says hello to Raha.
// nolint: wrapcheck
func (hh *Hello) Say(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString(fmt.Sprintf("Say %s from %s to Raha", hh.GreetingMessage, hh.Hostname))
}

// Register registers routes of hello handler on given group.
func (hh *Hello) Register(g fiber.Router) {
	g.Get("/", hh.Say)
}
