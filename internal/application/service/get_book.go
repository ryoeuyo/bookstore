package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ryoeuyo/test_go/internal/domain/book"
	"github.com/ryoeuyo/test_go/internal/infrastructure/repository/postgres/mapper"
)

func (s *Service) GetBook(ctx context.Context, id uuid.UUID) (*book.Book, error) {
	pgBook, err := s.Repository.GetBook(ctx, id)
	if err != nil {
		return nil, err
	}

	book := mapper.FromPostgresBook(pgBook)

	return &book, nil
}
