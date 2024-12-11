package service

import (
	"github.com/ryoeuyo/bookstore/internal/domain/book"
)

type BookService struct {
	Repo book.BookRepository
}

func NewBookService(r book.BookRepository) *BookService {
	return &BookService{
		Repo: r,
	}
}
