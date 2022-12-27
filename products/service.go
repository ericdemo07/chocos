package products

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	GetData(ctx context.Context)
}

type productServiceImpl struct {
	queryStore Store
}

func NewProductService(dbPool *pgxpool.Pool) Service {
	return &productServiceImpl{
		queryStore: NewStore(dbPool, 5000),
	}
}

func (productServiceImpl *productServiceImpl) GetData(ctx context.Context) {
	productServiceImpl.queryStore.getDataFromDB(ctx, "")
}
