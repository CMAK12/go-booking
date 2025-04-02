-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS hotels (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(250) NOT NULL,
    city VARCHAR(100) NOT NULL,
    description VARCHAR(1000),
    rating NUMERIC(3,1) DEFAULT 0
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO hotels (id, name, address, city, description, rating) VALUES
  (gen_random_uuid(), 'Hotel California', '123 Sunset Blvd', 'Los Angeles', 'A lovely hotel with a beautiful view.', 4.5),
  (gen_random_uuid(), 'The Grand Budapest Hotel', '456 Alpine St', 'Budapest', 'A charming hotel with a rich history.', 4.8),
  (gen_random_uuid(), 'The Shining Hotel', '789 Haunted Rd', 'Colorado', 'A spooky hotel with a mysterious past.', 3.9);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hotels;
-- +goose StatementEnd
