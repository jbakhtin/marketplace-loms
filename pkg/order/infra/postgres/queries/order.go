package queries

const (
	GetBySKU = `
		SELECT id, sku, name, price, created_at, updated_at FROM stocks
		WHERE sku = $1
		LIMIT 1
	`

	Reserve = `
		UPDATE SKU  quantity
	`

	ReserveCancel = `
		SKU string, quantity
	`

	ReserveRemove = `
		SKU quantity
	`
)
