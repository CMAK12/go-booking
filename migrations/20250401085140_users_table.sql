-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(319) UNIQUE NOT NULL,
  password VARCHAR(128) NOT NULL,
  role VARCHAR(50) NOT NULL CHECK (role IN ('guest', 'admin', 'manager')),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO users (id, name, email, password, role, created_at) VALUES
  ('bf61bd3f-3b0a-4282-9ec6-cc4441e03f62', 'John Doe', 'john.doe@example.com', 'hashed_password_1', 'admin', CURRENT_TIMESTAMP),
  ('5982a9e0-37b6-4318-9b14-fb14a165d523', 'Jane Smith', 'jane.smith@example.com', 'hashed_password_2', 'manager', CURRENT_TIMESTAMP),
  ('7e0bc3a1-3298-4d28-98be-67722ae75d79', 'Alice Johnson', 'alice.johnson@example.com', 'hashed_password_3', 'guest', CURRENT_TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
