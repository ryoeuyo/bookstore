-- +goose Up
ALTER TABLE books
DROP COLUMN date;

-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
ALTER TABLE books
ADD COLUMN date TIMESTAMP;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd
