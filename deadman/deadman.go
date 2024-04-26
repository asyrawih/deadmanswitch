package deadman

import (
	"kafka_compact/databases"
	"log/slog"
	"sync"
)

type Deadman struct {
	db *databases.Database
	mu *sync.Mutex
}

func NewDeadMan(db *databases.Database) *Deadman {
	return &Deadman{
		db: db,
		mu: new(sync.Mutex),
	}
}

func (dm *Deadman) Run(pair string) {
	slog.Info("spwan deadman watcher\t" + pair + "\n")
	if t, ok := dm.db.Get(pair); ok {
		<-t.C
		slog.Info(pair + " deadman timeout\n")
	}

	// cek if timer not stop
}
