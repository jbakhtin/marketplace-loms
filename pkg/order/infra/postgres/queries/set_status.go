package queries

const (
	SetStatusFailed = `
		UPDATE orders 
		SET status = 'failed'::order_statuses
		WHERE id = $1
		RETURNING id, user_id, status
	`

	SetStatusAwaitingPayment = `
		UPDATE orders 
		SET status = 'awaiting_payment'::order_statuses
		WHERE id = $1
		RETURNING id, user_id, status
	`
)
