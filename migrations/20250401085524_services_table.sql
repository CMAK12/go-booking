-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS services (
    id UUID PRIMARY KEY,
    room_id UUID REFERENCES rooms(id),
    name VARCHAR(255) NOT NULL,
    price NUMERIC(6,2) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO services (id, room_id, name, price) VALUES
  (gen_random_uuid(), '81b08cce-4ddf-45d7-a21c-2710653600b8', 'Room Service', 20.00),
  (gen_random_uuid(), '81b08cce-4ddf-45d7-a21c-2710653600b8', 'Laundry Service', 15.00),
  (gen_random_uuid(), '81b08cce-4ddf-45d7-a21c-2710653600b8', 'Spa Treatment', 100.00),
  (gen_random_uuid(), '81b08cce-4ddf-45d7-a21c-2710653600b8', 'Airport Shuttle', 50.00);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS services;
-- +goose StatementEnd
