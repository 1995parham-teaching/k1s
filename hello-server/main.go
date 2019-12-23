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

// ShutdownTimeout is a time for shuting down the echo server
const ShutdownTimeout = 5 * time.Second

func main() {
	fmt.Println("18.20 at Sep 07 2016 7:20 IR721")

	hostname, err := os.Hostname()
	if err != nil {
		logrus.Errorf("Cannot detect host name: %s", err)

		hostname = "parham"
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Say hello from %s to whom left me alone many years ago", hostname))
	})

	if err := e.Start(":1372"); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Server startup failed: %s", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("18.20 As always ... left me alone")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Printf("API Service failed on exit: %s", err)
	}
}
