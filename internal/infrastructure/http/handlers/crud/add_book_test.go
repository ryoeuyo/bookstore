package crud_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/crud"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/bookstore/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	svc := &service.BookService{
		Repo: mockRepo,
	}

	testBook := postgres.AddBookParams{
		Title:       "Test Book",
		Author:      "Test Author",
		Description: "Test Description",
		Genre:       "Fiction",
		Numberpages: 324,
	}

	testInvalidBook := postgres.AddBookParams{
		Numberpages: 0,
		Title:       "43242",
	}

	t.Run("successful add", func(t *testing.T) {
		reqBody, _ := json.Marshal(testBook)
		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(reqBody))
		defer req.Body.Close()

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		mockID := uuid.New()

		mockRepo.On("AddBook", context.Background(), testBook).Return(mockID, nil)

		router := gin.New()
		router.POST("/books", crud.AddBook(context.Background(), svc))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var resp = map[string]interface{}{}
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Error unmarshaling response body: %v", err)
		}

		assert.Equal(t, mockID.String(), resp["id"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("error adding book", func(t *testing.T) {
		reqBody, _ := json.Marshal(testInvalidBook)
		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		mockRepo.On("AddBook", context.Background(), testBook).Return(uuid.Nil, errors.New("failed to add book"))

		router := gin.New()
		router.POST("/books", crud.AddBook(context.Background(), svc))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		mockRepo.AssertExpectations(t)
	})
}
