package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ryoeuyo/test_go/internal/domain/book"
	"github.com/ryoeuyo/test_go/internal/infrastructure/repository/postgres"
)

func (s *Service) AddBook(ctx context.Context, book book.Book) (uuid.UUID, error) {
	if book.Author == "" || book.Title == "" {
		return uuid.Nil, errors.New("author and title must be set")
	}

	id, err := s.Repository.AddBook(ctx, postgres.AddBookParams{
		Title:       book.Title,
		Description: book.Description,
		Genre:       book.Genre,
		Author:      book.Author,
		Date:        pgtype.Timestamp{Time: book.Date, Valid: true},
		Numberpages: int32(book.NumberPages),
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
