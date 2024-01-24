package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	l := emptyLL[int]()
	assert.True(t, l.isEmptyLL())
}

func TestAppendElementIncrementSize(t *testing.T) {
	l := emptyLL[int]()
	assert.Equal(t, l.sizeLL(), 0)

	l.appendLL(10)

	assert.Equal(t, l.sizeLL(), 1)
}
 