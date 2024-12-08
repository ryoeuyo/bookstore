package service

import "github.com/ryoeuyo/bookstore/internal/domain/book"

type BookService struct {
	Repository book.BookRepository
}
