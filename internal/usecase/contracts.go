package usecase

import "github.com/andreyxaxa/Quotation-Book-service/internal/entity"

type (
	Quotes interface {
		AddQuote(entity.Quote) entity.Quote
		GetAllQuotes() []entity.Quote
		GetRandomQuote() (entity.Quote, error)
		GetQuotesByAuthor(string) []entity.Quote
		DeleteQuote(int) error
	}
)
