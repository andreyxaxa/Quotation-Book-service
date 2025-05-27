package persistent_test

import (
	"testing"

	"github.com/andreyxaxa/Quotation-Book-service/internal/entity"
	"github.com/andreyxaxa/Quotation-Book-service/internal/repo/persistent"
	"github.com/stretchr/testify/assert"
)

func TestAddAndGetAll(t *testing.T) {
	storage := persistent.New()

	q := entity.Quote{
		Author: "Confucius",
		Quote:  "Life is simple, but we insist on making it complicated.",
	}

	quote := storage.Add(q)

	assert.Equal(t, 1, quote.ID)
	assert.Equal(t, "Confucius", quote.Author)

	all := storage.GetAll()
	assert.Len(t, all, 1)
	assert.Equal(t, quote, all[0])
}

func TestGetRandom(t *testing.T) {
	storage := persistent.New()
	_, ok := storage.GetRandom()
	assert.False(t, ok)

	storage.Add(entity.Quote{
		Author: "Confucius",
		Quote:  "Life is simple, but we insist on making it complicated.",
	})

	q, ok := storage.GetRandom()
	assert.True(t, ok)
	assert.Equal(t, "Confucius", q.Author)
}

func TestGetByAuthor(t *testing.T) {
	storage := persistent.New()

	storage.Add(entity.Quote{
		Author: "A",
		Quote:  "Qu1",
	})
	storage.Add(entity.Quote{
		Author: "B",
		Quote:  "Qu2",
	})
	storage.Add(entity.Quote{
		Author: "A",
		Quote:  "Qu3",
	})

	res := storage.GetByAuthor("A")
	assert.Len(t, res, 2)
	assert.Equal(t, "Qu1", res[0].Quote)
	assert.Equal(t, "Qu3", res[1].Quote)
}

func TestDelete(t *testing.T) {
	storage := persistent.New()

	storage.Add(entity.Quote{
		Author: "A",
		Quote:  "Qu1",
	})
	storage.Add(entity.Quote{
		Author: "B",
		Quote:  "Qu2",
	})

	ok := storage.Delete(1)
	assert.True(t, ok)

	all := storage.GetAll()
	assert.Len(t, all, 1)
	assert.Equal(t, 2, all[0].ID)

	ok = storage.Delete(9999)
	assert.False(t, ok)
}
