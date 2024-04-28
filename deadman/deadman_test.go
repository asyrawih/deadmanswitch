package deadman

import (
	"kafka_compact/databases"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeadman_Run(t *testing.T) {
	d := databases.NewDatabase()
	dm := NewDeadMan(d)
	dm.db.Store("btcidr", 200)
	timers, ok := dm.db.Get("btcidr")
	assert.True(t, ok)

	assert.NotNil(t, timers)

	go dm.Run("btcidr")

	time.Sleep(1000 * time.Millisecond)
}

func BenchmarkRun(b *testing.B) {
	db := databases.NewDatabase()
	d := NewDeadMan(db)

	for i := 0; i < b.N; i++ {
		// Generate a unique key and duration for each iteration
		key := strconv.Itoa(i)

		// Call the Store method and measure its performance
		b.Run("Store-"+key, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				db.Store(key, 1000)
				go d.Run(key)
			}
		})
	}

	time.Sleep(1000 * time.Millisecond)

}
