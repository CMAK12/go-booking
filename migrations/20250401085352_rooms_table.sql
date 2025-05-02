-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rooms (
    id UUID PRIMARY KEY,
    hotel_id UUID REFERENCES hotels(id),
    type VARCHAR(100) NOT NULL,
    capacity INT NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    quantity INT NOT NULL DEFAULT 1
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_rooms_hotel_id ON rooms (hotel_id);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO rooms (id, hotel_id, type, capacity, price, quantity) VALUES
  ('7c5e05db-d822-4750-b9b8-881f2009f357', (SELECT id FROM hotels WHERE name = 'Hotel California'), 'Single', 1, 150.00, 10),
  ('e39c91e6-3674-4b29-871d-c87cb93a767f', (SELECT id FROM hotels WHERE name = 'Hotel California'), 'Double', 2, 200.00, 6),
  ('81b08cce-4ddf-45d7-a21c-2710653600b8', (SELECT id FROM hotels WHERE name = 'The Grand Budapest Hotel'), 'Suite', 4, 500.00, 4),
  (gen_random_uuid(), (SELECT id FROM hotels WHERE name = 'The Shining Hotel'), 'Deluxe', 2, 250.00, 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rooms;
-- +goose StatementEnd
