package databases

import (
	"log/slog"
	"sync"
	"time"
)

type (
	// databse
	// will hold the data by pairs
	Database struct {
		mu      *sync.Mutex
		tables  sync.Map
		Counter int64
	}
)

// create new databse
func NewDatabase() *Database {
	return &Database{}
}

// store data
func (db *Database) Store(key string, duration int64) {
	val, loaded := db.tables.Load(key)
	if loaded {
		// If the key exists, stop the existing timer
		timer := val.(*time.Timer)
		if !timer.Stop() {
			// If the timer has already expired, drain the channel
			// thats mean channel has expired not used need to create new one
			select {
			case <-timer.C:
				slog.Info("timer expired before resetting")
				db.tables.Delete(key)
			default:
				slog.Info("delete key")
				db.tables.Delete(key)
			}
		}
	}

	db.tables.Store(key, time.NewTimer(time.Duration(duration)*time.Millisecond))

}

// delete data from tables
func (db *Database) Delete(key string) {
	slog.Info("delete key: " + key)
	db.tables.Delete(key)
}

// get data from tables
func (db *Database) Get(key string) (*time.Timer, bool) {
	val, ok := db.tables.Load(key)
	if val == nil {
		return nil, false
	}
	return val.(*time.Timer), ok

}

func (db *Database) Reset(pair string, duration int64) {
	if value, ok := db.tables.Load(pair); ok {
		value.(*time.Timer).Reset(time.Duration(duration) * time.Millisecond)
	}
}
