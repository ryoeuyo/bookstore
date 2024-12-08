package service

import "github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"

func IsValidBook(book *postgres.AddBookParams) bool {
	if book.Author == "" || book.Title == "" || book.Genre == "" || book.Numberpages == 0 {
		return false
	}

	return true
}
