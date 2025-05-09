package postgres

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
)

type OrderStorage struct {
}

func NewOrderStorage() (OrderStorage, error) {
	return OrderStorage{}, nil
}

func (o OrderStorage) Reserve(ctx context.Context, SKU string, quantity int) error {
	//TODO implement me
	return nil
}

func (o OrderStorage) ReserveCancel(ctx context.Context, SKU string, quantity int) error {
	//TODO implement me
	return nil
}

func (o OrderStorage) ReserveRemove(ctx context.Context, SKU string, quantity int) error {
	//TODO implement me
	return nil
}

func (o OrderStorage) GetBySKU(ctx context.Context, SKU string) ([]domain.StockItem, error) {
	//TODO implement me
	return make([]domain.StockItem, 0), nil
}

func (o OrderStorage) Create(ctx context.Context, order domain.Order) (domain.Order, error) {
	//TODO implement me
	return domain.Order{}, nil
}

func (o OrderStorage) SetStatus(ctx context.Context, ID int, status string) (domain.Order, error) {
	//TODO implement me
	return domain.Order{}, nil
}

func (o OrderStorage) GetByID(ctx context.Context, ID int) (domain.Order, error) {
	//TODO implement me
	return domain.Order{}, nil
}
