-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_items (
    id bigserial NOT NULL PRIMARY KEY,
    sku int NOT NULL,
    count int,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_items;
-- +goose StatementEnd