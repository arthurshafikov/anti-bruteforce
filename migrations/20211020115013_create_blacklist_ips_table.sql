-- +goose Up
-- +goose StatementBegin
CREATE TABLE blacklist_ips (
    id serial primary key,
    subnet INET NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS blacklist_ips;
-- +goose StatementEnd
