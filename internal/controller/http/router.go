package http

import (
	v1 "github.com/andreyxaxa/Quotation-Book-service/internal/controller/http/v1"
	"github.com/andreyxaxa/Quotation-Book-service/internal/usecase"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, q usecase.Quotes) {
	//apiV1Group := r.PathPrefix("/v1").Subrouter()
	{
		v1.NewQuotesRoutes(r, q)
	}
}
