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

func TestAllBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	svc := &service.BookService{
		Repo: mockRepo,
	}

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	defer req.Body.Close()

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
	router.GET("/books", crud.AllBooks(context.Background(), svc))

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
}
