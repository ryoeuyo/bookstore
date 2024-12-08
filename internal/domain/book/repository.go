package book

import (
	"context"

	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

type BookRepository interface {
	AllBooks(ctx context.Context) ([]postgres.Book, error)
	GetBook(ctx context.Context, id uuid.UUID) (postgres.Book, error)
	AddBook(ctx context.Context, params postgres.AddBookParams) (uuid.UUID, error)
	DeleteBook(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}
