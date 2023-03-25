package data

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDateComponents(t *testing.T) {
	date := time.Date(1965, time.April, 23, 12, 2, 0, 0, time.UTC)
	got := NewDateComponents(date)

	want := DateComponents{
		Year:  1965,
		Month: 4,
		Day:   23,
	}

	assert.Equal(t, *got, want)
}
