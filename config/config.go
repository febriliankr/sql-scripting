package config

import (
	"fmt"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	DatabaseURL    string `env:"DATABASE_URL"`
	DatabaseDriver string `env:"DATABASE_DRIVER" envDefault:"postgres"`
}

func InitConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg

}
