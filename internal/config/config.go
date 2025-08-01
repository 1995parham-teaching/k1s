package config

import (
	"log"
	"strings"

	"github.com/1995parham-teaching/k1s/internal/logger"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

const Namespace = "K1S_"

type Config struct {
	Server Server        `koanf:"server"`
	Logger logger.Config `koanf:"logger"`
}

type Server struct {
	Port            int    `koanf:"port"`
	GreetingMessage string `koanf:"greeting"`
}

func Init(configFileName string) Config {
	cfg := new(Config)
	kn := koanf.New(".")

	err := kn.Load(Default(), nil)
	if err != nil {
		log.Printf("error loading defaults: %s", err)
	}
	// load configuration from file
	err := kn.Load(file.Provider(configFileName), yaml.Parser())
	if err != nil {
		log.Printf("error loading %s: %s", configFileName, err)
	}

	// load environment variables
	err := kn.Load(env.Provider(Namespace, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, Namespace)), "_", ".")
	}), nil)
	if err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	err := kn.Unmarshal("", cfg)
	if err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	log.Printf("following configuration is loaded:\n%+v", *cfg)

	return *cfg
}
