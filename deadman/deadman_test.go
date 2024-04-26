package deadman

import (
	"kafka_compact/databases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeadman_Run(t *testing.T) {
	d := databases.NewDatabase()
	dm := NewDeadMan(d)
	dm.db.Store("btcidr", 200)
	timers := dm.db.Get("btcidr")
	assert.NotNil(t, timers)

	go dm.Run("btcidr")

	time.Sleep(1000 * time.Millisecond)
}
