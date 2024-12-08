package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/bookstore/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)

	s := &service.BookService{
		Repository: mockRepo,
	}

	testBook := postgres.AddBookParams{
		Title:       "Test Book",
		Author:      "Test Author",
		Description: "Test Desc",
		Genre:       "Test Genre",
		Numberpages: 343,
	}

	mockRepo.On("AddBook", context.Background(), testBook).Return(uuid.New(), nil)

	id, err := s.AddBook(context.Background(), testBook)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

	mockRepo.AssertExpectations(t)

	testInvalidBook := postgres.AddBookParams{
		Numberpages: 0,
		Title:       "43242",
	}

	id, err = s.AddBook(context.Background(), testInvalidBook)

	assert.Error(t, err)
	assert.Equal(t, uuid.Nil, id)

	mockRepo.AssertExpectations(t)
}
