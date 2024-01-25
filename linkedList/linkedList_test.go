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

func TestConsElementIncrementSize(t *testing.T) {
	l := emptyLL[int]()
	assert.Equal(t, l.sizeLL(), 0)
	l.consLL(10)
	assert.Equal(t, l.sizeLL(), 1)
}

func TestContainsElement(t *testing.T) {
	l := emptyLL[int]()
	l.appendLL(19)
	l.appendLL(33)
	l.appendLL(19)
	l.appendLL(1)
	l.appendLL(22)
	l.appendLL(2)
	l.appendLL(3)
	l.consLL(5)

	assert.True(t, l.containsLL(5))
	assert.True(t, l.containsLL(3))
}
