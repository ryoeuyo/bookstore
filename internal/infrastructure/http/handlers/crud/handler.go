package crud

import (
	"github.com/go-playground/validator/v10"
	"github.com/ryoeuyo/bookstore/internal/application/service"
)

type BookHandler struct {
	Svc   *service.BookService
	Valid *validator.Validate
}

func NewBookHandler(s *service.BookService, v *validator.Validate) *BookHandler {
	return &BookHandler{
		Svc:   s,
		Valid: v,
	}
}
