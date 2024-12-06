package book

import "github.com/google/uuid"

type BookRepository interface {
	FetchAll() (*[]Book, error)
	FetchByUUID(uuid uuid.UUID) (*Book, error)
	Create(newBook *Book) (uuid.UUID, error)
	Update(UUID uuid.UUID) error
	Delete(UUID uuid.UUID) error
}
