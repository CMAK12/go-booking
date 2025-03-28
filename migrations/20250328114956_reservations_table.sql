-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reservations (
		id UUID PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100),
		date DATE NOT NULL,
		guest_quantity INT NOT NULL,
		city VARCHAR(100) NOT NULL
	);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reservations;
-- +goose StatementEnd
