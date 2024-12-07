-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: AllBooks :many
SELECT * FROM books;

-- name: AddBook :one
INSERT INTO books (
    title, description,
    genre, author,
    date, numberPages
) VALUES (
    $1, $2,
    $3, $4,
    $5, $6
)
RETURNING id;

-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING id;
