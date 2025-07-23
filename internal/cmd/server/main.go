package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1995parham-teaching/k1s/internal/config"
	"github.com/1995parham-teaching/k1s/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// ShutdownTimeout is a time for shutting down the echo server.
const ShutdownTimeout = 5 * time.Second

func main(cfg config.Config, logger *zap.Logger) {
	f := fiber.New()

	hh := handler.NewHello(cfg.Server.GreetingMessage, logger.Named("http.hello"))
	hh.Register(f.Group(""))

	h := handler.NewHealth()
	h.Register(f.Group(""))

	err := f.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		logger.Fatal("Server startup failed", zap.Error(err))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err := f.Shutdown()
	if err != nil {
		logger.Error("API Service failed on exit", zap.Error(err))
	}
}

func Register(root *cobra.Command, cfg config.Config, logger *zap.Logger) {
	// nolint: exhaustruct
	root.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run the hello server",
		Run: func(_ *cobra.Command, _ []string) {
			main(cfg, logger)
		},
	})
}
