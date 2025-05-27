package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andreyxaxa/Quotation-Book-service/internal/entity"
	"github.com/gorilla/mux"
)

func (h *V1) add(w http.ResponseWriter, r *http.Request) {
	var q entity.Quote

	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	quote := h.q.AddQuote(q)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quote)
}

func (h *V1) getAll(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var quotes []entity.Quote

	if author != "" {
		quotes = h.q.GetQuotesByAuthor(author)
	} else {
		quotes = h.q.GetAllQuotes()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *V1) getRandom(w http.ResponseWriter, r *http.Request) {
	q, err := h.q.GetRandomQuote()
	if err != nil {
		http.Error(w, "no quotes found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func (h *V1) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.q.DeleteQuote(id)
	if err != nil {
		http.Error(w, "quote not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
