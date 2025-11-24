package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
)

type Order struct {
	Id        uint             `json:"id,omitempty" db:"id"`
	Status    string           `json:"status,omitempty" db:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" db:"updated_at"`
}

func (o *Order) ToModel() models.Order {
	return models.Order{
		ID:     o.Id,
		Status: o.Status,
		Items:  make([]models.OrderItem, 0),
	}
}
