-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER set_timestamp_trigger_orders_items
    BEFORE UPDATE ON order_items
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER set_timestamp_trigger_orders_items;
-- +goose StatementEnd
