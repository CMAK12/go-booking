-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS booking_services (
    booking_id UUID REFERENCES bookings(id) ON DELETE CASCADE,
    service_id UUID REFERENCES services(id) ON DELETE CASCADE,
    PRIMARY KEY (booking_id, service_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS booking_services;
-- +goose StatementEnd
