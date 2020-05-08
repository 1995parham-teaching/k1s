package config

import (
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/sirupsen/logrus"
)

const Namespace = "K1S_"

type Config struct {
	Server Server `koanf:"server"`
}

type Server struct {
	Port            int    `koanf:"port"`
	GreetingMessage string `koanf:"greeting"`
}

func Init(configFileName string) Config {
	cfg := new(Config)
	k := koanf.New(".")

	if err := k.Load(Default(), nil); err != nil {
		logrus.Errorf("error loading defaults: %s", err)
	}
	// load configuration from file
	if err := k.Load(file.Provider(configFileName), yaml.Parser()); err != nil {
		logrus.Errorf("error loading %s: %s", configFileName, err)
	}

	// load environment variables
	if err := k.Load(env.Provider(Namespace, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, Namespace)), "_", ".", -1)
	}), nil); err != nil {
		logrus.Errorf("error loading environment variables: %s", err)
	}

	if err := k.Unmarshal("", cfg); err != nil {
		logrus.Fatalf("error unmarshalling config: %s", err)
	}

	logrus.Infof("following configuration is loaded:\n%+v", *cfg)

	return *cfg
}
