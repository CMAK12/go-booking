-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS discounts (
    id UUID PRIMARY KEY,
    hotel_id UUID REFERENCES hotels(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    active BOOLEAN DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO discounts (id, hotel_id, name, amount, active) VALUES
  (gen_random_uuid(), (SELECT id FROM hotels WHERE name = 'Hotel California'), 'Winter Discount', 20.00, TRUE),
  (gen_random_uuid(), (SELECT id FROM hotels WHERE name = 'Hotel California'), 'Summer Special', 15.00, TRUE),
  (gen_random_uuid(), (SELECT id FROM hotels WHERE name = 'The Grand Budapest Hotel'), 'Loyalty Program', 10.00, TRUE),
  (gen_random_uuid(), (SELECT id FROM hotels WHERE name = 'The Shining Hotel'), 'Early Bird', 25.00, FALSE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS discounts;
-- +goose StatementEnd
