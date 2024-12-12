package book

import (
	"context"

	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

type Repository interface {
	AllBooks(ctx context.Context) ([]postgres.Book, error)
	GetBook(ctx context.Context, id uuid.UUID) (postgres.Book, error)
	AddBook(ctx context.Context, params postgres.AddBookParams) (uuid.UUID, error)
	DeleteBook(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	UpdateBook(ctx context.Context, arg postgres.UpdateBookParams) (uuid.UUID, error)
	UpdateTitleBook(ctx context.Context, arg postgres.UpdateTitleBookParams) (uuid.UUID, error)
	UpdateNumberPagesBook(ctx context.Context, arg postgres.UpdateNumberPagesBookParams) (uuid.UUID, error)
	UpdateGenreBook(ctx context.Context, arg postgres.UpdateGenreBookParams) (uuid.UUID, error)
	UpdateDescriptionBook(ctx context.Context, arg postgres.UpdateDescriptionBookParams) (uuid.UUID, error)
	UpdateAuthorBook(ctx context.Context, arg postgres.UpdateAuthorBookParams) (uuid.UUID, error)
}
