package crud_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/crud"
	"github.com/ryoeuyo/bookstore/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	svc := &service.BookService{
		Repository: mockRepo,
	}

	t.Run("successful delete", func(t *testing.T) {
		mockID := uuid.New()
		req, _ := http.NewRequest(http.MethodDelete, "/books/"+mockID.String(), nil)

		rr := httptest.NewRecorder()

		mockRepo.On("DeleteBook", context.Background(), mockID).Return(mockID, nil)

		router := gin.New()
		router.DELETE("/books/:id", crud.DeleteBook(context.Background(), svc))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var resp = map[string]interface{}{}
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Error unmarshaling response body: %v", err)
		}

		assert.Equal(t, mockID.String(), resp["id"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("error delete book", func(t *testing.T) {
		mockInvalidID := "invalid_id"
		req, _ := http.NewRequest(http.MethodDelete, "/books/"+mockInvalidID, nil)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := gin.New()
		router.DELETE("/books/:id", crud.DeleteBook(context.Background(), svc))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		mockRepo.AssertExpectations(t)
	})
}
