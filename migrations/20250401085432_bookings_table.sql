-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bookings (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    room_id UUID REFERENCES rooms(id) ON DELETE CASCADE,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    status VARCHAR(30) NOT NULL CHECK (status IN ('pending', 'confirmed', 'cancelled'))
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE bookings
ADD CONSTRAINT check_dates CHECK (start_date < end_date);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO bookings (id, user_id, room_id, start_date, end_date, status) VALUES
    (gen_random_uuid(), 'bf61bd3f-3b0a-4282-9ec6-cc4441e03f62', '7c5e05db-d822-4750-b9b8-881f2009f357', '2025-04-01 11:00:00', '2025-04-05 10:00:00', 'confirmed'),
    (gen_random_uuid(), 'bf61bd3f-3b0a-4282-9ec6-cc4441e03f62', 'e39c91e6-3674-4b29-871d-c87cb93a767f', '2025-04-10 11:00:00', '2025-04-15 10:00:00', 'pending'),
    (gen_random_uuid(), '5982a9e0-37b6-4318-9b14-fb14a165d523', '81b08cce-4ddf-45d7-a21c-2710653600b8', '2025-04-20 11:00:00', '2025-04-25 10:00:00', 'cancelled');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bookings;
-- +goose StatementEnd
