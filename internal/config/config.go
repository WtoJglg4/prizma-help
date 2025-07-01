package config

import (
	"os"
)

type Config struct {
	HTTPAddr string
}

func Load() Config {
	cfg := Config{
		HTTPAddr: ":5000",
	}
	httpAddr, ok := os.LookupEnv("HTTP_ADDR")
	if ok {
		cfg.HTTPAddr = httpAddr
	}
	return cfg
}
