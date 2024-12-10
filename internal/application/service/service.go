package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/ryoeuyo/bookstore/internal/domain/book"
	"github.com/ryoeuyo/bookstore/internal/shared/logger/validate"
)

type BookService struct {
	Repo     book.BookRepository
	Validate *validator.Validate
}

func NewBookService(r book.BookRepository) *BookService {
	v := validator.New()
	v.RegisterValidation("notzero", validate.IsNotZero)
	v.RegisterValidation("notempty", validate.IsNotEmpty)

	return &BookService{
		Repo:     r,
		Validate: v,
	}
}
