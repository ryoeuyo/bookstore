package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func (s *BookService) GetBook(ctx context.Context, id uuid.UUID) (*postgres.Book, error) {
	book, err := s.Repo.GetBook(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New(ErrNotExists)
		}

		return nil, err
	}

	return &book, nil
}
