-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
 		id UUID PRIMARY KEY,
 		username VARCHAR(100) UNIQUE NOT NULL,
 		password VARCHAR(500) NOT NULL,
 		email VARCHAR(100) UNIQUE NOT NULL,
 		role VARCHAR(100) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
