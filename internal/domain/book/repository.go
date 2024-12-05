package book

import "github.com/google/uuid"

type BookRepository interface {
	FetchAll() (error, *[]Book)
	FetchByUUID(uuid uuid.UUID) (error, *Book)
	Create(NewBook *Book) error
}
