-- +goose Up
-- +goose StatementBegin
CREATE TYPE reserved_item_statuses as enum('RESERVED', 'COMMITTED', 'RELEASED');

CREATE TABLE reserved_items (
    id bigserial NOT NULL PRIMARY KEY,
    order_id bigint NOT NULL,
    sku int NOT NULL,
    reserved int NOT NULL,
    status reserved_item_statuses NOT NULL DEFAULT 'RESERVED'::reserved_item_statuses,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reserved_items;
DROP TYPE reserved_item_statuses;
-- +goose StatementEnd
