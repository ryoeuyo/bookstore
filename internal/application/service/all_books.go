package service

import (
	"context"
	"errors"

	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func (s *BookService) AllBooks(ctx context.Context) (*[]postgres.Book, error) {
	books, err := s.Repo.AllBooks(ctx)
	if err != nil {
		return nil, err
	}

	if books == nil {
		return nil, errors.New(ErrNotFound)
	}

	return &books, nil
}
