-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: AllBooks :many
SELECT * FROM books;

-- name: CreateBook :one
INSERT INTO books (
    title, description,
    genre, author,
    date, quantityOnStock,
    numberPages, price
) VALUES (
    $1, $2,
    $3, $4,
    $5, $6,
    $7, $8
)
RETURNING id;

-- name: UpdatePrice :one
UPDATE books
set price = $2
WHERE id = $1
RETURNING id;

-- name: UpdateQuantity :one
UPDATE books
set quantityOnStock = $2
WHERE id = $1
RETURNING id;

-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING id;
