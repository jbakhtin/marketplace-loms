-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER set_timestamp_trigger_reserved_items
    BEFORE UPDATE ON reserved_items
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER set_timestamp_trigger_reserved_items;
-- +goose StatementEnd

