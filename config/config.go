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
	Port int
}

type DatabaseConfigurations struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
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
			Port: viper.GetInt("server.port"),
		},
		database: DatabaseConfigurations{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			DBName:   viper.GetString("database.dbname"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
		},
	}
}

func (conf Configurations) ServerConfiguration() ServerConfigurations {
	return conf.server
}

func (conf Configurations) DatabaseConfig() DatabaseConfigurations {
	return conf.database
}
