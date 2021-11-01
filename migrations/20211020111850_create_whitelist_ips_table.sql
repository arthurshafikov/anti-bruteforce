-- +goose Up
-- +goose StatementBegin
CREATE TABLE whitelist_ips (
    id serial primary key,
    subnet INET NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS whitelist_ips;
-- +goose StatementEnd
