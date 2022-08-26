package config

import (
	"github.com/1995parham-teaching/k1s/internal/logger"
	"github.com/knadh/koanf/providers/structs"
)

const Port = 1378

func Default() *structs.Structs {
	return structs.Provider(Config{
		Logger: logger.Config{
			Level: "debug",
		},
		Server: Server{
			Port:            Port,
			GreetingMessage: "hello with default",
		},
	}, "koanf")
}
