package botvk

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/jmoiron/sqlx"
)

func Start(dbURL string, apiVK *api.VK, gpoudID int) error {
	return nil
}

func newDB(dbURL string) (*sqlx.DB, error) {
	db,err:= sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}