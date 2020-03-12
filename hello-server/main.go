/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 04-05-2019
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ShutdownTimeout is a time for shutting down the echo server
const ShutdownTimeout = 5 * time.Second

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		logrus.Errorf("Cannot detect host name: %s", err)

		hostname = "parham"
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Say hello from %s to Raha", hostname))
	})
	e.GET("/healthz", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	if err := e.Start(":1378"); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Server startup failed: %s", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Printf("API Service failed on exit: %s", err)
	}
}
