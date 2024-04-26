package databases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	d := NewDatabase()
	assert.NotNil(t, d)
}

func TestDatabase_Store(t *testing.T) {
	d := NewDatabase()
	d.Store("btcidr", 100)
	d.Store("btcidr", 100)
}

func TestDatabase_Delete(t *testing.T) {
	d := NewDatabase()
	d.Store("btcidr", 100)
}

func TestDatabase_Get(t *testing.T) {
	d := NewDatabase()
	d.Store("btcidr", 200)
	target, b := d.Get("btcidr")
	if !b {
		t.Fail()
	}
	t.Log(target)
}
