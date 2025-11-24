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

type Repository struct {
	db *sql.DB
}

func NewStockRepository(db *sql.DB) (Repository, error) {
	return Repository{
		db: db,
	}, nil
}

func (r *Repository) Reserve(
	ctx context.Context,
	orderId uint,
	reservationItems []models.ReservationItem,
) (models.Reservation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Reservation{}, errors.Wrap(err, "repository: reserve: begin transaction")
	}
	defer tx.Rollback()

	for _, reservationItem := range reservationItems {
		var available uint64
		err = tx.QueryRowContext(ctx, queries.GetStockItemForUpdate, reservationItem.SKU).Scan(&available)
		if err != nil {
			return models.Reservation{}, errors.Wrap(err, "repository: reserve: get stock items for update")
		}

		if available < reservationItem.Reserved {
			return models.Reservation{}, errors.Wrap(errors.New("not enough stocks"), "repository: reserve")
		}
	}

	reservedItems := make([]entities.ReservationItem, 0, len(reservationItems))
	for i := range reservationItems {
		var reservedItem entities.ReservationItem
		err = tx.QueryRowContext(ctx, queries.InsertReservationItem,
			orderId,
			reservationItems[i].SKU,
			reservationItems[i].Reserved,
		).Scan(
			&reservedItem.ID,
			&reservedItem.OrderID,
			&reservedItem.SKU,
			&reservedItem.Reserved,
			&reservedItem.Status,
			&reservedItem.CreatedAt,
			&reservedItem.UpdatedAt,
		)
		if err != nil {
			return models.Reservation{}, errors.Wrap(err, "repository: reserve: insert reservation item")
		}

		_, err = tx.ExecContext(ctx, queries.UpdateStockItem,
			reservationItems[i].Reserved,
			reservationItems[i].SKU,
		)
		if err != nil {
			return models.Reservation{}, errors.Wrap(err, "repository: reserve: update stock item")
		}

		reservedItems = append(reservedItems, reservedItem)
	}

	err = tx.Commit()
	if err != nil {
		return models.Reservation{}, errors.Wrap(err, "reserve")
	}

	reservationModel := models.Reservation{
		OrderId:          orderId,
		ReservationItems: make([]models.ReservationItem, 0, len(reservedItems)),
	}

	for _, item := range reservedItems {
		reservationModel.ReservationItems = append(reservationModel.ReservationItems, item.ToModel())
	}

	return reservationModel, nil
}

func (r *Repository) GetReservationByOrderID(ctx context.Context, orderId uint) (models.Reservation, error) {
	query, err := r.db.QueryContext(ctx, queries.GetReservationByOrderId, orderId)
	if err != nil {
		return models.Reservation{}, errors.Wrap(err, "get reservation by order id")
	}
	defer query.Close()

	reservations := make([]entities.ReservationItem, 0)

	for query.Next() {
		var reservation entities.ReservationItem
		err = query.Scan(
			&reservation.ID,
			&reservation.OrderID,
			&reservation.SKU,
			&reservation.Reserved,
			&reservation.Status,
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
		)
		if err != nil {
			return models.Reservation{}, errors.Wrap(err, "get reservation by order id")
		}

		reservations = append(reservations, reservation)
	}

	if len(reservations) <= 0 {
		return models.Reservation{}, domain.NotFoundError
	}

	reservationModel := models.Reservation{
		OrderId:          uint(reservations[0].OrderID),
		ReservationItems: make([]models.ReservationItem, len(reservations)),
	}

	for _, reservation := range reservations {
		reservationModel.ReservationItems = append(reservationModel.ReservationItems, reservation.ToModel())
	}

	return reservationModel, nil
}

func (r *Repository) ReserveCancel(ctx context.Context, orderId uint) (models.Reservation, error) {
	//TODO implement me
	return models.Reservation{}, nil
}

func (r *Repository) ReserveRemove(ctx context.Context, orderId uint) (models.Reservation, error) {
	//TODO implement me
	return models.Reservation{}, nil
}

func (r *Repository) GetStockItemBySKU(ctx context.Context, SKU int32) (models.StockItem, error) {
	var product entities.StockItem

	err := r.db.QueryRowContext(ctx, queries.GetBySKU, SKU).Scan(
		&product.ID,
		&product.SKU,
		&product.Available,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.StockItem{}, domain.NotFoundError
		}
		return models.StockItem{}, errors.Wrap(err, "failed to get product by sku")
	}

	return product.ToModel(), nil
}
