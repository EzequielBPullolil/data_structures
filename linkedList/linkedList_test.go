package linkedlist

import (
	"fmt"
	"slices"
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

func TestToListIsInOrdConsLL(t *testing.T) {
	l := emptyLL[int]()
	elements := [6]int{10, 9, 3, 2, 5, 34}
	for i := range elements {
		l.consLL(elements[i])
	}
	assert.Equal(t, elements[:], l.toList())
}

func TestToListIsInOrdAppendLL(t *testing.T) {
	l := emptyLL[int]()
	elements := []int{10, 9, 3, 2, 5, 34}
	for i := range elements {
		l.appendLL(elements[i])
	}
	slices.Reverse(elements)
	assert.Equal(t, elements[:], l.toList())
}

func TestInsertAt(t *testing.T) {
	l := emptyLL[int]()
	elements := []int{10, 9, 3, 2, 4, 5, 34}
	for i := range elements {
		if i == 4 {
			continue
		}
		l.consLL(elements[i])
	}
	l.insertAt(4, 4)

	fmt.Println(l.toList())
	assert.Equal(t, elements[4], 4)
}

func TestPopElement(t *testing.T) {
	l := emptyLL[int]()
	elements := [6]int{10, 9, 3, 2, 5, 34}
	for i := range elements {
		l.consLL(elements[i])
	}

	assert.Equal(t, elements[:], l.toList())

	l.popLL()
	assert.Equal(t, elements[1:], l.toList())
}
