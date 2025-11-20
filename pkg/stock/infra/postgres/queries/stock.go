package queries

const (
	GetBySKU = `
		SELECT id, sku, available, created_at, updated_at FROM stock_items
		WHERE sku = $1
		LIMIT 1
	`

	Reserve = `
		WITH updated AS (
			UPDATE stock_items 
				SET available = available - $2
			WHERE sku = $1 AND available >= $2
			RETURNING sku
		)
		INSERT INTO reserved_items (sku, reserved, status)
		SELECT $1, $2, 'RESERVED'::reserved_item_statuses
		WHERE EXISTS (SELECT 1 FROM updated)
	`

	ReserveCancel = `
		SKU string, quantity
	`

	ReserveRemove = `
		SKU quantity
	`
)
