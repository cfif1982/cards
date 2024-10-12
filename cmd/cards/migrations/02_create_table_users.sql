-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
  uuid UUID PRIMARY KEY,
  name      TEXT,
	last_name   TEXT,
	email       TEXT,
	telephone TEXT,
	login TEXT,
	password TEXT
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd