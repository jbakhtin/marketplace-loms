-- +goose Up
-- +goose StatementBegin
CREATE TYPE order_statuses as enum('new', 'awaiting_payment', 'failed', 'payed', 'cancelled');

CREATE TABLE orders (
    id bigserial NOT NULL PRIMARY KEY,
    sku int NOT NULL,
    status order_statuses NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
DROP TYPE order_statuses;
-- +goose StatementEnd
