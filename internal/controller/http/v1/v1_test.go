package v1_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/andreyxaxa/Quotation-Book-service/internal/controller/http/v1"
	"github.com/andreyxaxa/Quotation-Book-service/internal/entity"
	"github.com/andreyxaxa/Quotation-Book-service/internal/repo/persistent"
	"github.com/andreyxaxa/Quotation-Book-service/internal/usecase/quotes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	storage := persistent.New()
	uc := quotes.New(storage)
	r := mux.NewRouter()
	v1.NewQuotesRoutes(r, uc)

	return r
}

func TestAddQuote(t *testing.T) {
	router := setupRouter()

	body := `{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}`
	req := httptest.NewRequest("POST", "/quotes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var q entity.Quote
	err := json.NewDecoder(w.Body).Decode(&q)
	assert.NoError(t, err)
	assert.Equal(t, "Confucius", q.Author)
}

func TestGetAllQuotes(t *testing.T) {
	router := setupRouter()

	// Add quote
	body := `{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}`
	req := httptest.NewRequest("POST", "/quotes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Get all quotes
	req = httptest.NewRequest("GET", "/quotes", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var quotes []entity.Quote
	err := json.NewDecoder(w.Body).Decode(&quotes)
	assert.NoError(t, err)
	assert.Len(t, quotes, 1)
}

func TestDeleteQuote(t *testing.T) {
	router := setupRouter()

	// Add quote
	body := `{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}`
	req := httptest.NewRequest("POST", "/quotes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Delete quote
	req = httptest.NewRequest("DELETE", "/quotes/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
