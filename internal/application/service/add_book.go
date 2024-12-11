package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func (s *BookService) AddBook(ctx context.Context, book postgres.AddBookParams) (uuid.UUID, error) {
	id, err := s.Repo.AddBook(ctx, book)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
