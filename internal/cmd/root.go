package cmd

import (
	"os"

	"github.com/1995parham/k1s/internal/cmd/server"
	"github.com/1995parham/k1s/internal/config"
	"github.com/1995parham/k1s/internal/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// ExitFailure is returned status in case of failure.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	root := new(cobra.Command)
	root.Use = "hello-server"
	root.Short = "Say hello to kubernetes"

	cfg := config.Init("config.yaml")

	logger := logger.New(cfg.Logger)

	server.Register(root, cfg, logger)

	if err := root.Execute(); err != nil {
		logger.Error("failed to execute root command", zap.Error(err))

		os.Exit(ExitFailure)
	}
}
