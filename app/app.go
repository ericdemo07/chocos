package app

import (
	"example/hello/config"
	"example/hello/database"
	"example/hello/products"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependency struct {
	ProductService products.Service
	Config         config.Configurations
}

func Init(conf config.Configurations) *Dependency {
	dbPool := database.Connection(conf.DatabaseConfig())

	return initDependencies(dbPool, conf)
}

func initDependencies(dbPool *pgxpool.Pool, conf config.Configurations) *Dependency {
	ps := products.NewProductService(dbPool)

	return &Dependency{
		ProductService: ps,
		Config:         conf,
	}
}
