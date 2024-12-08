-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE books (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    author VARCHAR(256) NOT NULL,
    date TIMESTAMP NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    title VARCHAR(256) NOT NULL,
    description VARCHAR NOT NULL,
    genre VARCHAR(256) NOT NULL,
    numberPages INTEGER NOT NULL
);

-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
DROP TABLE books;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd
