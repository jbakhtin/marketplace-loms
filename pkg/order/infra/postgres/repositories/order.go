package repositories

import (
	"context"
	"database/sql"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
	"github.com/jbakhtin/marketplace-loms/pkg/order/infra/postgres/entities"
	"github.com/jbakhtin/marketplace-loms/pkg/order/infra/postgres/queries"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) (OrderRepository, error) {
	return OrderRepository{
		db: db,
	}, nil
}

func (o OrderRepository) Create(ctx context.Context, userID uint64, items []models.OrderItem) (models.Order, error) {
	var order entities.Order
	err := o.db.QueryRowContext(ctx, queries.CreateOrder, userID).Scan(&order.Id, &order.UserID, &order.Status)
	if err != nil {
		return models.Order{}, err
	}

	orderItems := make([]entities.OrderItem, len(items))
	for i, item := range items {
		err = o.db.QueryRowContext(ctx, queries.CreateOrderItem, order.Id, item.SKU, item.Count).Scan(
			&orderItems[i].ID,
			&orderItems[i].OrderID,
			&orderItems[i].SKU,
			&orderItems[i].Quantity,
			&orderItems[i].CreatedAt,
			&orderItems[i].UpdatedAt,
		)
		if err != nil {
			return models.Order{}, err
		}
	}

	orderModel := order.ToModel()

	for _, orderItem := range orderItems {
		orderModel.Items = append(orderModel.Items, orderItem.ToModel())
	}

	return orderModel, nil
}

func (o OrderRepository) SetStatus(ctx context.Context, ID int64, status string) (models.Order, error) {
	//TODO implement me
	return models.Order{}, nil
}

func (o OrderRepository) GetByID(ctx context.Context, ID int64) (models.Order, error) {
	//TODO implement me
	return models.Order{}, nil
}
