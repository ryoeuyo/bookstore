package book

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	UUID            uuid.UUID `json:"uuid"`
	Author          string    `json:"author"`
	Date            time.Time `json:"date"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Genre           string    `json:"genre"`
	NumberPages     int       `json:"numberPages"`
	Price           float64   `json:"price"`
	QuantityOnStock int       `json:"quantityOnStock"`
}

func NewBook(
	UUID uuid.UUID,
	Title,
	Description,
	Genre,
	Author string,
	Date,
	CreatedAt,
	UpdatedAt time.Time,
	QuantityOnStock,
	NumberPages int,
	Price float64,
) *Book {
	return &Book{
		UUID:            UUID,
		Title:           Title,
		Description:     Description,
		Genre:           Genre,
		Author:          Author,
		Date:            Date,
		CreatedAt:       CreatedAt,
		UpdatedAt:       UpdatedAt,
		QuantityOnStock: QuantityOnStock,
		NumberPages:     NumberPages,
		Price:           Price,
	}
}
