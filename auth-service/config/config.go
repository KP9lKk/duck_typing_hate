package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App  App
		GRPC GRPC
		Log  Log
		RDB  RDB
	}
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}
	GRPC struct {
		Port string `env:"GRPC_PORT,required"`
	}
	RDB struct {
		Url      string `env:"RDB_URL,required"`
		Password string `env:"RDB_PASSWORD,required"`
		Db       int    `env:"RDB_DB,required"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
