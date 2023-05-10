package model

type Quote struct {
	ID int 	`db:"id"`
	Text string	`db:"quote_text"`
	PeerID int	`db:"peer_id"`
}

type Quotes []Quote