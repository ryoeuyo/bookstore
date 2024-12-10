-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: AllBooks :many
SELECT * FROM books;

-- name: AddBook :one
INSERT INTO books (
    title, description,
    genre, author,
    numberPages
) VALUES (
    $1, $2,
    $3, $4,
    $5
)
RETURNING id;

-- name: UpdateTitleBook :one
UPDATE books
SET title = $2,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: UpdateAuthorBook :one
UPDATE books
SET author = $2,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: UpdateDescriptionBook :one
UPDATE books
SET description = $2,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: UpdateGenreBook :one
UPDATE books
SET genre = $2,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: UpdateNumberPagesBook :one
UPDATE books
SET numberPages = $2,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: UpdateBook :one
UPDATE books
set title = $2,
    description = $3,
    genre = $4,
    author = $5,
    numberPages = $6,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING id;
