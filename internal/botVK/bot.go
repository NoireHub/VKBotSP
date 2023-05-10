package botvk

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NoireHub/VKBotSP/internal/store/sqlstore"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Start(dbURL string, apiVK *api.VK, groupID int) error {
	db, err := newDB(dbURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	s, err := NewServer(apiVK, store, groupID)
	if err != nil {
		return err
	}
	graceFulShutdown(s)

	s.logger.Info("Bot is up and ready")
	if err := s.longPoll.Run(); err != nil {
		return err
	}

	return nil
}

func graceFulShutdown(s *server) {
	c := make(chan os.Signal,1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		s.longPoll.Shutdown()
		s.logger.Fatal("interrupt")
	}()

}

func newDB(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
