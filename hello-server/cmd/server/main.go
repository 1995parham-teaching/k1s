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

package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1995parham/k1s/hello-server/config"

	"github.com/1995parham/k1s/hello-server/handler"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ShutdownTimeout is a time for shutting down the echo server.
const ShutdownTimeout = 5 * time.Second

func main(cfg config.Config) {
	e := echo.New()

	hh := handler.NewHello(cfg.Server.GreetingMessage)
	hh.Register(e.Group(""))

	h := handler.NewHealth()
	h.Register(e.Group(""))

	if err := e.Start(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil && errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("Server startup failed: %s", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logrus.Errorf("API Service failed on exit: %s", err)
	}
}

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run the hello server",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	})
}
