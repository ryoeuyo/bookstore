package service

import (
	"context"

	"github.com/ryoeuyo/test_go/internal/domain/book"
	"github.com/ryoeuyo/test_go/internal/infrastructure/repository/postgres/mapper"
)

type Service struct {
	Repository book.BookRepository
}

func (s Service) AllBooks(ctx context.Context) ([]book.Book, error) {
	pgBooks, err := s.Repository.AllBooks(ctx)
	if err != nil {
		return nil, err
	}

	var books []book.Book

	// Mapping database model to book.Book
	for _, b := range pgBooks {
		book := mapper.FromPostgresBook(b)
		books = append(books, book)
	}

	return books, nil
}
