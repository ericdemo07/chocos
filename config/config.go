package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	port int
}

func Load() Config {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return Config{
		port: extractIntOrDefault("APP_PORT", 3030),
	}
}

func (conf Config) Port() int {
	return conf.port
}
