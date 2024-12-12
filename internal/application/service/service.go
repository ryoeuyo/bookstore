package service

import (
	"github.com/ryoeuyo/bookstore/internal/domain/book"
)

type BookService struct {
	Repo book.Repository
}

func NewBookService(r book.Repository) *BookService {
	return &BookService{
		Repo: r,
	}
}
