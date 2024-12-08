package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func (s *BookService) AddBook(ctx context.Context, book postgres.AddBookParams) (uuid.UUID, error) {
	if !IsValidBook(&book) {
		return uuid.Nil, errors.New(ErrInvalidBook)
	}

	id, err := s.Repository.AddBook(ctx, book)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
