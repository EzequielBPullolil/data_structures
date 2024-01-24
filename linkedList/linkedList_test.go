package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	l := emptyLL[int]()
	assert.True(t, l.isEmptyLL())
}
