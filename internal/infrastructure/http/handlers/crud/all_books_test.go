package crud_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/crud"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/etc/validate"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/bookstore/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAllBook(t *testing.T) {
	t.Run("all books", func(t *testing.T) {
		mockRepo := new(mocks.BookRepository)

		v := validator.New()
		v.RegisterValidation("notzero", validate.IsNotZero)
		v.RegisterValidation("notempty", validate.IsNotEmpty)

		svc := &service.BookService{
			Repo: mockRepo,
		}

		handler := crud.NewBookHandler(svc, v)

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)

		rr := httptest.NewRecorder()
		testOutputBooks := []postgres.Book{
			{
				ID:          uuid.New(),
				Createdat:   time.Now(),
				Updatedat:   time.Now(),
				Title:       "Test Book",
				Author:      "Test Author",
				Description: "Test Description",
				Genre:       "Fiction",
				Numberpages: 324,
			},
			{
				ID:          uuid.New(),
				Createdat:   time.Now(),
				Updatedat:   time.Now(),
				Title:       "Test Book",
				Author:      "Test Author",
				Description: "Test Description",
				Genre:       "Fiction",
				Numberpages: 324,
			},
			{
				ID:          uuid.New(),
				Createdat:   time.Now(),
				Updatedat:   time.Now(),
				Title:       "Test Book",
				Author:      "Test Author",
				Description: "Test Description",
				Genre:       "Fiction",
				Numberpages: 324,
			},
		}
		mockRepo.On("AllBooks", context.Background()).Return(testOutputBooks, nil)

		router := gin.New()
		router.GET("/books", handler.AllBooks(context.Background()))

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var resp []postgres.Book
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Error unmarshaling response body: %v", err)
		}

		for i := range resp {
			// Zeroing fields
			resp[i].Createdat = time.Time{}
			resp[i].Updatedat = time.Time{}
			testOutputBooks[i].Createdat = time.Time{}
			testOutputBooks[i].Updatedat = time.Time{}

			assert.Equal(t, testOutputBooks[i], resp[i])
		}

		mockRepo.AssertExpectations(t)
	})
}
