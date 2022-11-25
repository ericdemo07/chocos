package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	server      ServerConfigurations
	database    DatabaseConfigurations
	examplePath string
	exampleVar  string
}

type ServerConfigurations struct {
	port int
}

type DatabaseConfigurations struct {
	name     string
	user     string
	password string
}

func Load() Configurations {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return Configurations{
		server: ServerConfigurations{
			port: viper.GetInt("server.port"),
		},
	}
}

func (conf Configurations) Port() int {
	return conf.server.port
}
