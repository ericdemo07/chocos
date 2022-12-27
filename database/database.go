package database

import (
	"context"
	"example/hello/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

const databaseUrlFormat = "postgres://%s:%s@%s:%s/%s"

func Connection(conf config.DatabaseConfigurations) *pgxpool.Pool {

	databaseUrl := fmt.Sprintf(databaseUrlFormat, conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)

	dbPool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	//defer dbPool.Close()

	var greeting string

	err = dbPool.QueryRow(context.Background(), "select 'Hello, world 1!'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	return dbPool
}
