-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_items (
    id bigserial NOT NULL PRIMARY KEY,
    order_id bigint NOT NULL,
    sku int NOT NULL,
    quantity int,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp,
    FOREIGN KEY (order_id) REFERENCES orders (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_items;
-- +goose StatementEnd