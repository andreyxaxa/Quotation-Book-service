package v1

import (
	"net/http"

	"github.com/andreyxaxa/Quotation-Book-service/internal/usecase"
	"github.com/gorilla/mux"
)

func NewQuotesRoutes(apiV1Group *mux.Router, q usecase.Quotes) {
	r := &V1{q: q}

	quotesGroup := apiV1Group.PathPrefix("/quotes").Subrouter()

	quotesGroup.HandleFunc("", r.add).Methods(http.MethodPost)
	quotesGroup.HandleFunc("", r.getAll).Methods(http.MethodGet)
	quotesGroup.HandleFunc("/random", r.getRandom).Methods(http.MethodGet)
	quotesGroup.HandleFunc("/{id:[0-9]+}", r.delete).Methods(http.MethodDelete)
}
