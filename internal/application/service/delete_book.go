package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteBook(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	id, err := s.Repository.DeleteBook(ctx, id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
