package apiserver

import "github.com/bilyalovdenis/testserver/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

func NewConfig() *Config{
	return &Config{
		BindAddr: ":4000",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}