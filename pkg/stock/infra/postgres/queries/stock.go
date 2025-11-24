package queries

const (
	GetBySKU = `
		SELECT id, sku, available, created_at, updated_at FROM stock_items
		WHERE sku = $1
		LIMIT 1
	`

	GetReservationByOrderId = `
		SELECT id, order_id, sku, reserved, status, created_at, updated_at
		FROM reserved_items
		WHERE order_id = $1
	`
)
