package products

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

type Store interface {
	getDataFromDB(ctx context.Context, id string)
}

type productsDB struct {
	dbPool       *pgxpool.Pool
	queryTimeout time.Duration
}

func NewStore(dbPool *pgxpool.Pool, queryTimeout time.Duration) Store {

	return &productsDB{
		dbPool:       dbPool,
		queryTimeout: queryTimeout,
	}
}

func (productsDB *productsDB) getDataFromDB(ctx context.Context, id string) {
	var greeting string

	err := productsDB.dbPool.QueryRow(context.Background(), "select 'Hello, world 2!'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
