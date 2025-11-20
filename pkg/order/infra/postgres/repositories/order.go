package repositories

import (
	"context"
	"database/sql"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) (OrderRepository, error) {
	return OrderRepository{
		db: db,
	}, nil
}

func (o OrderRepository) Create(ctx context.Context, order entity.Order) (entity.Order, error) {
	//TODO implement me
	return entity.Order{}, nil
}

func (o OrderRepository) SetStatus(ctx context.Context, ID int, status string) (entity.Order, error) {
	//TODO implement me
	return entity.Order{}, nil
}

func (o OrderRepository) GetByID(ctx context.Context, ID int) (entity.Order, error) {
	//TODO implement me
	return entity.Order{}, nil
}
