package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Hello handles the hello messages.
type Hello struct {
	Hostname string
}

// NewHello creates a new instance of hello handler.
func NewHello() *Hello {
	hostname, err := os.Hostname()
	if err != nil {
		logrus.Errorf("cannot detect host name: %s", err)

		hostname = "parham"
	}

	return &Hello{
		Hostname: hostname,
	}
}

// Say says hello to Raha.
func (hh *Hello) Say(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Say hello from %s to Raha", hh.Hostname))
}

// Register registers routes of hello handler on given group.
func (hh *Hello) Register(g *echo.Group) {
	g.GET("/", hh.Say)
}
