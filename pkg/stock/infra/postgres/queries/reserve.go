package queries

var (
	GetStockItemForUpdate = `
		SELECT available
		FROM stock_items 
		WHERE sku = $1
		FOR UPDATE
	`

	UpdateStockItem = `
		UPDATE stock_items
		SET available = available - $1
		WHERE sku = $2
	`

	InsertReservationItem = `
		INSERT INTO reserved_items (order_id, sku, reserved)
		VALUES ($1, $2, $3)
		RETURNING id, order_id, sku, reserved, status, created_at, updated_at
	`
)
