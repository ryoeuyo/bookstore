package crud_test

import (
	"bytes"
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
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/bookstore/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpdateBook(t *testing.T) {

	t.Run("success update", func(t *testing.T) {
		mockRepo := new(mocks.BookRepository)

		v := validator.New()
		v.RegisterValidation("notzero", validate.IsNotZero)
		v.RegisterValidation("notempty", validate.IsNotEmpty)

		svc := &service.BookService{
			Repo: mockRepo,
		}

		handler := crud.NewBookHandler(svc, v)
		mockID := uuid.New()
		testBook := postgres.UpdateBookParams{
			ID:          mockID,
			Title:       "Test Book",
			Author:      "Test Author",
			Description: "Test Description",
			Genre:       "Fiction",
			Numberpages: 324,
		}

		reqBody, _ := json.Marshal(testBook)
		req, _ := http.NewRequest(http.MethodPut, "/books", bytes.NewReader(reqBody))
		defer req.Body.Close()

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		mockRepo.On("UpdateBook", context.Background(), testBook).Return(mockID, nil)

		router := gin.New()
		router.PUT("/books", handler.UpdateBook(context.Background()))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var resp = map[string]interface{}{}
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Error unmarshaling response body: %v", err)
		}

		assert.Equal(t, mockID.String(), resp["id"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("with invalid book data", func(t *testing.T) {
		mockRepo := new(mocks.BookRepository)

		v := validator.New()
		v.RegisterValidation("notzero", validate.IsNotZero)
		v.RegisterValidation("notempty", validate.IsNotEmpty)

		svc := &service.BookService{
			Repo: mockRepo,
		}

		handler := crud.NewBookHandler(svc, v)
		mockID := uuid.New()
		testBook := postgres.UpdateBookParams{
			ID:          mockID,
			Title:       "",
			Numberpages: -23,
		}

		reqBody, _ := json.Marshal(testBook)
		req, _ := http.NewRequest(http.MethodPut, "/books", bytes.NewReader(reqBody))

		rr := httptest.NewRecorder()

		router := gin.New()
		router.PUT("/books", handler.UpdateBook(context.Background()))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestUpdateFieldBook(t *testing.T) {

}
