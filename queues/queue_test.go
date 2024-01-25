package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := emptyQ[string]()
	assert.True(t, q.isEmptyQ())

	q.enqueue("ezequiel")

	assert.Equal(t, "ezequiel", q.firstQ())
}
func createQueueSuject() (Queue[string], []string) {
	q := emptyQ[string]()
	elements := []string{"ezequiel", "jose", "sofia", "lucia", "juan", "alguno"}
	for i := range elements {
		q.enqueue(elements[i])
	}
	return q, elements
}
func TestFirstQ(t *testing.T) {
	q, _ := createQueueSuject()
	assert.Equal(t, "ezequiel", q.firstQ())
}

func TestDequeueQ(t *testing.T) {
	q, _ := createQueueSuject()
	assert.Equal(t, "ezequiel", q.firstQ())
	q.dequeue()
	assert.Equal(t, "jose", q.firstQ())

	assert.Equal(t, len(q.bq), 0)
}

func TestFullQueue(t *testing.T) {
	q, elements := createQueueSuject()
	for i := range elements {
		assert.Equal(t, q.firstQ(), elements[i])
		q.dequeue()
	}

	q.enqueue("element")
	assert.Equal(t, len(q.bq), 0)
	q.enqueue("element2")
	assert.Equal(t, len(q.bq), 1)
}
