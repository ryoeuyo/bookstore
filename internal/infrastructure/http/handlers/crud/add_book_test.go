package crud_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryoeuyo/bookstore/internal/etc/validate"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/crud"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/bookstore/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)

	v := validator.New()
	v.RegisterValidation("notzero", validate.IsNotZero)
	v.RegisterValidation("notempty", validate.IsNotEmpty)

	svc := &service.BookService{
		Repo: mockRepo,
	}

	handler := crud.NewBookHandler(svc, v)

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
		Description: "",
		Genre:       "",
	}

	t.Run("successful add", func(t *testing.T) {
		reqBody, _ := json.Marshal(testBook)
		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(reqBody))

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		mockID := uuid.New()

		mockRepo.On("AddBook", context.Background(), testBook).Return(mockID, nil)

		router := gin.New()
		router.POST("/books", handler.AddBook(context.Background()))
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

		mockRepo.On("AddBook", context.Background(), testInvalidBook).Return(uuid.Nil, errors.New("failed to add book"))

		router := gin.New()
		router.POST("/books", handler.AddBook(context.Background()))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		mockRepo.AssertExpectations(t)
	})
}
