package queries

const (
	CreateOrder = `
		INSERT INTO orders (status)
		VALUES (DEFAULT)
		RETURNING id, status
	`

	CreateOrderItem = `
		INSERT INTO order_items (order_id, sku, quantity)
		VALUES ($1, $2, $3)
		RETURNING id, order_id, sku, quantity, created_at, updated_at
	`
)
