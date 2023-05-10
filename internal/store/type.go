package store

import "github.com/NoireHub/VKBotSP/internal/model"

type Store interface {
	Quote() QuoteRepository
}

type QuoteRepository interface {
	Create(*model.Quote) error
	FindById(int) ([]model.Quote, error)
}