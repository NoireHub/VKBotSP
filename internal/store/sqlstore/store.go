package sqlstore

import (
	"github.com/NoireHub/VKBotSP/internal/store"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db             *sqlx.DB
	QuoteRepository *QuoteRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,

	}
}

func (s *Store) Quote() store.QuoteRepository {
	if s.QuoteRepository != nil {
		return s.QuoteRepository
	}

	s.QuoteRepository = &QuoteRepository{
		store: s,
	}

	return s.QuoteRepository
}