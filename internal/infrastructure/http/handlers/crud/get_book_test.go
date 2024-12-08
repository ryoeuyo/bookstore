package crud_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/crud"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/bookstore/internal/mocks"
	"github.com/stretchr/testify/assert"
)


func TestGetBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	svc := &service.BookService{
		Repository: mockRepo,
	}

	t.Run("successful get", func(t *testing.T) {
		mockID := uuid.New()
		testOutputBook := postgres.Book{
			ID:          mockID,
			Createdat:   time.Now(),
			Updatedat:   time.Now(),
			Title:       "Test Book",
			Author:      "Test Author",
			Description: "Test Description",
			Genre:       "Fiction",
			Numberpages: 323,
		}

		req, _ := http.NewRequest(http.MethodGet, "/books/"+mockID.String(), nil)

		rr := httptest.NewRecorder()

		mockRepo.On("GetBook", context.Background(), mockID).Return(testOutputBook, nil)

		router := gin.New()
		router.GET("/books/:id", crud.GetBook(context.Background(), svc))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var resp postgres.Book
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Error unmarshaling response body: %v", err)
		}

		resp.Createdat = time.Time{}
		resp.Updatedat = time.Time{}
		testOutputBook.Createdat = time.Time{}
		testOutputBook.Updatedat = time.Time{}

		assert.Equal(t, testOutputBook, resp)

		mockRepo.AssertExpectations(t)
	})

	t.Run("error get book", func(t *testing.T) {
		mockInvalidID := "invalid_id"
		req, _ := http.NewRequest(http.MethodGet, "/books/"+mockInvalidID, nil)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := gin.New()
		router.GET("/books/:id", crud.DeleteBook(context.Background(), svc))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		mockRepo.AssertExpectations(t)
	})
}
