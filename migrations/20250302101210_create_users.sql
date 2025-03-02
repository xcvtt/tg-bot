-- +goose Up
-- +goose StatementBegin
create table users (
    id bigint,
    name text,
    hp int
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
