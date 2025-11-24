package queries

const (
	CreateOrder = `
		INSERT INTO orders (user_id, status)
		VALUES ($1, DEFAULT)
		RETURNING id, user_id, status
	`

	CreateOrderItem = `
		INSERT INTO order_items (order_id, sku, quantity)
		VALUES ($1, $2, $3)
		RETURNING id, order_id, sku, quantity, created_at, updated_at
	`
)
