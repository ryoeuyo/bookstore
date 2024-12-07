package mapper

import (
	"github.com/ryoeuyo/test_go/internal/domain/book"
	"github.com/ryoeuyo/test_go/internal/infrastructure/repository/postgres"
)

func FromPostgresBook(pgBook postgres.Book) book.Book {
	return book.Book{
		UUID:        pgBook.ID,
		Author:      pgBook.Author,
		Date:        pgBook.Date.Time, // pgtype.Timestamp to time.Time
		CreatedAt:   pgBook.Createdat.Time,
		UpdatedAt:   pgBook.Updatedat.Time,
		Title:       pgBook.Title,
		Description: pgBook.Description,
		Genre:       pgBook.Genre,
		NumberPages: int(pgBook.Numberpages),
	}
}
