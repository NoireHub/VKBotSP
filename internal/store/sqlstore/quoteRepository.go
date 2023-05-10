package sqlstore

import (
	"errors"
	"strings"

	"github.com/NoireHub/VKBotSP/internal/model"
)

type QuoteRepository struct {
	store *Store
}

func (q *QuoteRepository) Create(quote *model.Quote) error {

	_, err := q.store.db.Exec(
		"INSERT INTO quotes(quote_text, peer_id) VALUES($1,$2)",
		quote.Text,
		quote.PeerID,
	)
	if err != nil {
		if strings.Contains(err.Error(),"pq: duplicate key") {
			return errors.New("такая цитата уже существует")
		}

		return err
	}

	return nil
}

func (q *QuoteRepository) FindById(peerID int) ([]model.Quote, error) {
	quote:= []model.Quote{}
	rows, err:= q.store.db.Queryx(
		"SELECT * FROM quotes WHERE peer_id = $1",
		peerID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tempQuote model.Quote

		rows.StructScan(&tempQuote)
		quote = append(quote, tempQuote)
	}

	return quote, nil
}