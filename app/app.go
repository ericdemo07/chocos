package app

import (
	"example/hello/config"
	"example/hello/database"
)

func Init(conf config.Configurations) {
	database.DatabaseConnection(conf.DatabaseConfig())
}
