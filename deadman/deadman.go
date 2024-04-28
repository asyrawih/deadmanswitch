package deadman

import (
	"kafka_compact/databases"
	"log/slog"
)

type Deadman struct {
	db *databases.Database
}

func NewDeadMan(db *databases.Database) *Deadman {
	return &Deadman{
		db: db,
	}
}

func (dm *Deadman) Run(pair string) {
	slog.Info("spwan deadman watcher\t" + pair + "\n")
	if t, ok := dm.db.Get(pair); ok {
		<-t.C
		slog.Info(pair + " deadman timeout\n")
		dm.db.Delete(pair)
	}
}
