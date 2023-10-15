-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id serial PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  email VARCHAR(150) NOT null UNIQUE,
  password VARCHAR(80) NOT NULL,
  is_active BOOLEAN NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
