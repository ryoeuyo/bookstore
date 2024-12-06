// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const allBooks = `-- name: AllBooks :many
SELECT id, author, date, createdat, updatedat, title, description, genre, numberpages, price, quantityonstock FROM books
`

func (q *Queries) AllBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.Query(ctx, allBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Date,
			&i.Createdat,
			&i.Updatedat,
			&i.Title,
			&i.Description,
			&i.Genre,
			&i.Numberpages,
			&i.Price,
			&i.Quantityonstock,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createBook = `-- name: CreateBook :one
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
RETURNING id
`

type CreateBookParams struct {
	Title           string
	Description     string
	Genre           string
	Author          string
	Date            pgtype.Timestamp
	Quantityonstock int32
	Numberpages     int32
	Price           pgtype.Numeric
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createBook,
		arg.Title,
		arg.Description,
		arg.Genre,
		arg.Author,
		arg.Date,
		arg.Quantityonstock,
		arg.Numberpages,
		arg.Price,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteBook = `-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteBook(ctx context.Context, id pgtype.UUID) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, deleteBook, id)
	err := row.Scan(&id)
	return id, err
}

const getBook = `-- name: GetBook :one
SELECT id, author, date, createdat, updatedat, title, description, genre, numberpages, price, quantityonstock FROM books
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id pgtype.UUID) (Book, error) {
	row := q.db.QueryRow(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Date,
		&i.Createdat,
		&i.Updatedat,
		&i.Title,
		&i.Description,
		&i.Genre,
		&i.Numberpages,
		&i.Price,
		&i.Quantityonstock,
	)
	return i, err
}

const updatePrice = `-- name: UpdatePrice :one
UPDATE books
set price = $2
WHERE id = $1
RETURNING id
`

type UpdatePriceParams struct {
	ID    pgtype.UUID
	Price pgtype.Numeric
}

func (q *Queries) UpdatePrice(ctx context.Context, arg UpdatePriceParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, updatePrice, arg.ID, arg.Price)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const updateQuantity = `-- name: UpdateQuantity :one
UPDATE books
set quantityOnStock = $2
WHERE id = $1
RETURNING id
`

type UpdateQuantityParams struct {
	ID              pgtype.UUID
	Quantityonstock int32
}

func (q *Queries) UpdateQuantity(ctx context.Context, arg UpdateQuantityParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, updateQuantity, arg.ID, arg.Quantityonstock)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}