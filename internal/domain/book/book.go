package book

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	UUID        uuid.UUID `json:"uuid"`
	Author      string    `json:"author"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	NumberPages int       `json:"numberPages"`
}
