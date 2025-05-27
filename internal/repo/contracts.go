package repo

import "github.com/andreyxaxa/Quotation-Book-service/internal/entity"

type (
	QuotesRepo interface {
		Add(entity.Quote) entity.Quote
		GetAll() []entity.Quote
		GetRandom() (entity.Quote, bool)
		GetByAuthor(string) []entity.Quote
		Delete(int) bool
	}
)
