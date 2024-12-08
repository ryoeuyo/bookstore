package service

import (
	"context"

	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func (s *BookService) AllBooks(ctx context.Context) (*[]postgres.Book, error) {
	books, err := s.Repository.AllBooks(ctx)
	if err != nil {
		return nil, err
	}

	return &books, nil
}
