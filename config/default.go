package config

import (
	"github.com/knadh/koanf/providers/structs"
)

const Port = 1378

func Default() *structs.Structs {
	return structs.Provider(Config{
		Server: Server{
			Port:            Port,
			GreetingMessage: "hello with default",
		},
	}, "koanf")
}
