package crud_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryoeuyo/bookstore/internal/etc/validate"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/crud"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)

	v := validator.New()
	v.RegisterValidation("notzero", validate.IsNotZero)
	v.RegisterValidation("notempty", validate.IsNotEmpty)

	svc := &service.BookService{
		Repo: mockRepo,
	}

	handler := crud.NewBookHandler(svc, v)

	t.Run("successful delete", func(t *testing.T) {
		mockID := uuid.New()
		req, _ := http.NewRequest(http.MethodDelete, "/books/"+mockID.String(), nil)

		rr := httptest.NewRecorder()

		mockRepo.On("DeleteBook", context.Background(), mockID).Return(mockID, nil)

		router := gin.New()
		router.DELETE("/books/:id", handler.DeleteBook(context.Background()))

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
		router.DELETE("/books/:id", handler.DeleteBook(context.Background()))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		mockRepo.AssertExpectations(t)
	})
}
