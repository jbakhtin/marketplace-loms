-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS unique_index_stock_items_sku ON stock_items(sku);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS unique_index_stock_items_sku;
-- +goose StatementEnd