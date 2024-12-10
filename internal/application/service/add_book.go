package service

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func (s *BookService) AddBook(ctx context.Context, book postgres.AddBookParams) (uuid.UUID, error) {
	err := s.Validate.Struct(book)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return uuid.Nil, errors.New(ErrValidation)
		}

		return uuid.Nil, errors.New(ErrInvalidBook)
	}

	id, err := s.Repo.AddBook(ctx, book)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
