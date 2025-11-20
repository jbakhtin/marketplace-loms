package repositories

import (
	"context"
	"database/sql"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain/models"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/infra/postgres/entities"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/infra/postgres/queries"
	"github.com/pkg/errors"
)

type StockRepository struct {
	db *sql.DB
}

func NewStockRepository(db *sql.DB) (StockRepository, error) {
	return StockRepository{
		db: db,
	}, nil
}

func (o *StockRepository) Reserve(ctx context.Context, SKU int32, qty uint16) error {
	//TODO implement me
	return nil
}

func (o *StockRepository) ReserveCancel(ctx context.Context, SKU int32, qty uint16) error {
	//TODO implement me
	return nil
}

func (o *StockRepository) ReserveRemove(ctx context.Context, SKU int32, qty uint16) error {
	//TODO implement me
	return nil
}

func (o *StockRepository) GetBySKU(ctx context.Context, SKU int32) (models.StockItem, error) {
	var product entities.StockItem

	err := o.db.QueryRowContext(ctx, queries.GetBySKU, SKU).Scan(
		&product.ID,
		&product.SKU,
		&product.Available,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.StockItem{}, domain.NotFound
		}
		return models.StockItem{}, errors.Wrap(err, "failed to get product by sku")
	}

	return product.ToModel(), nil
}
