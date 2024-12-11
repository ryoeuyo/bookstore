package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func (s *BookService) UpdateBook(ctx context.Context, updBook postgres.UpdateBookParams) (uuid.UUID, error) {
	id, err := s.Repo.UpdateBook(ctx, updBook)
	if err != nil {
		if err == pgx.ErrNoRows {
			return uuid.Nil, errors.New(ErrNotExists)
		}

		return uuid.Nil, err
	}

	return id, nil
}

func (s *BookService) UpdateFieldBook(ctx context.Context, id uuid.UUID, field, value string) (uuid.UUID, error) {
	var err error
	var idOut uuid.UUID

	switch field {
	case "title":
		idOut, err = s.Repo.UpdateTitleBook(ctx, postgres.UpdateTitleBookParams{
			ID:    id,
			Title: value,
		})
	case "description":
		idOut, err = s.Repo.UpdateDescriptionBook(ctx, postgres.UpdateDescriptionBookParams{
			ID:          id,
			Description: value,
		})
	case "author":
		idOut, err = s.Repo.UpdateAuthorBook(ctx, postgres.UpdateAuthorBookParams{
			ID:     id,
			Author: value,
		})
	case "genre":
		idOut, err = s.Repo.UpdateGenreBook(ctx, postgres.UpdateGenreBookParams{
			ID:    id,
			Genre: value,
		})
	case "numberpages":
		number, err := strconv.ParseInt(value, 32, 10)
		if err != nil {
			return uuid.Nil, errors.New(ErrInvalidField)
		}

		idOut, err = s.Repo.UpdateNumberPagesBook(ctx, postgres.UpdateNumberPagesBookParams{
			ID:          id,
			Numberpages: int32(number),
		})
	default:
		return uuid.Nil, errors.New(ErrInvalidField)
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			return uuid.Nil, errors.New(ErrNotExists)
		}

		return uuid.Nil, err
	}

	return idOut, nil
}
