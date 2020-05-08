package config

import "github.com/knadh/koanf/providers/confmap"

const Port = 1378

func Default() *confmap.Confmap {
	return confmap.Provider(map[string]interface{}{
		"server.port": Port,
	},
		".",
	)
}
