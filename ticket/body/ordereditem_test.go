package body

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedListItemSegment_Value(t *testing.T) {
	item := NewOrderedListItemSegment(1, 10)
	assert.Equal(t, "10", item.Value())
}
