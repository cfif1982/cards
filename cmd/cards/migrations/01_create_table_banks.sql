-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banks(
  uuid UUID PRIMARY KEY,
  name      TEXT,
	address   TEXT,
	bik       int,
	telephone TEXT
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banks;
-- +goose StatementEnd
