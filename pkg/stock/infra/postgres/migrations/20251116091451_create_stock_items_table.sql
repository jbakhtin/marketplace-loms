-- +goose Up
-- +goose StatementBegin
CREATE TABLE stock_items (
     id bigserial NOT NULL PRIMARY KEY,
     sku int NOT NULL,
     available int,
     created_at timestamp NOT NULL DEFAULT now(),
     updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock_items;
-- +goose StatementEnd
