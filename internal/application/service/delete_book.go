package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *BookService) DeleteBook(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	id, err := s.Repo.DeleteBook(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return uuid.Nil, errors.New(ErrNotExists)
		}

		return uuid.Nil, err
	}

	return id, nil
}
