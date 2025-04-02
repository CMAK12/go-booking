-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS services (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO services (id, name, price) VALUES
  (gen_random_uuid(), 'Room Service', 20.00),
  (gen_random_uuid(), 'Laundry Service', 15.00),
  (gen_random_uuid(), 'Spa Treatment', 100.00),
  (gen_random_uuid(), 'Airport Shuttle', 50.00);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS services;
-- +goose StatementEnd
