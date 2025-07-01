package config

import (
	"fmt"
	"os"
	"sync/atomic"

	"github.com/pelletier/go-toml/v2"
)

var (
	cfg                Config
	loadedSuccessfully atomic.Bool
)

type Http struct {
	Port string `toml:"port"`
}

type Config struct {
	Env      string `toml:"env"`
	LogLevel string `toml:"log_level"`
	Http     Http   `toml:"http"`
}

func GetConfig(path string) (*Config, error) {
	if !loadedSuccessfully.Load() {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("can't read config file: %s\n", err)
			return nil, err
		}

		if err := toml.Unmarshal(data, &cfg); err != nil {
			fmt.Printf("can't parse config file: %s\n", err)
			return nil, err
		}
		loadedSuccessfully.Store(true)
	}
	return &cfg, nil
}
