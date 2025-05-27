package quotes

import (
	"errors"

	"github.com/andreyxaxa/Quotation-Book-service/internal/entity"
	"github.com/andreyxaxa/Quotation-Book-service/internal/repo"
)

type UseCase struct {
	repo repo.QuotesRepo
}

func New(r repo.QuotesRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) AddQuote(q entity.Quote) entity.Quote {
	return uc.repo.Add(q)
}

func (uc *UseCase) GetAllQuotes() []entity.Quote {
	return uc.repo.GetAll()
}

func (uc *UseCase) GetRandomQuote() (entity.Quote, error) {
	q, ok := uc.repo.GetRandom()
	if !ok {
		return entity.Quote{}, errors.New("no quotes available")
	}

	return q, nil
}

func (uc *UseCase) GetQuotesByAuthor(author string) []entity.Quote {
	return uc.repo.GetByAuthor(author)
}

func (uc *UseCase) DeleteQuote(id int) error {
	ok := uc.repo.Delete(id)
	if !ok {
		return errors.New("quote not found")
	}

	return nil
}
