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
	Ttile           string    `json:"title"`
	Genre           string    `json:"genre"`
	NumberPages     int64     `json:"numberPages"`
	Price           float64   `json:"price"`
	QuantityOnStock int64     `json:"quantityOnStock"`
}

func NewBook(
	UUID uuid.UUID,
	Title,
	Genre,
	Author string,
	Date,
	CreatedAt,
	UpdatedAt time.Time,
	QuantityOnStock,
	NumberPages int64,
	Price float64,
) *Book {
	return &Book{
		UUID:            UUID,
		Ttile:           Title,
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
