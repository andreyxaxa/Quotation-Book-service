package persistent

import (
	"math/rand"
	"sync"

	"github.com/andreyxaxa/Quotation-Book-service/internal/entity"
)

type QuoteStorage struct {
	mu     sync.Mutex
	quotes []entity.Quote
	ID     int
}

func New() *QuoteStorage {
	return &QuoteStorage{
		quotes: []entity.Quote{},
		ID:     1,
	}
}

func (s *QuoteStorage) Add(quote entity.Quote) entity.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	quote.ID = s.ID
	s.ID++

	s.quotes = append(s.quotes, quote)

	return quote
}

func (s *QuoteStorage) GetAll() []entity.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	var all []entity.Quote
	all = append(all, s.quotes...)

	return all
}

func (s *QuoteStorage) GetRandom() (entity.Quote, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.quotes) == 0 {
		return entity.Quote{}, false
	}

	return s.quotes[rand.Intn(len(s.quotes))], true
}

func (s *QuoteStorage) GetByAuthor(author string) []entity.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	var filtered []entity.Quote

	for _, q := range s.quotes {
		if q.Author == author {
			filtered = append(filtered, q)
		}
	}

	return filtered
}

func (s *QuoteStorage) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, q := range s.quotes {
		if q.ID == id {
			s.quotes = append(s.quotes[:i], s.quotes[i+1:]...)
			return true
		}
	}

	return false
}
