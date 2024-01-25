package queues

/*
Dado un Queue con los campos FQ y BQ
1. Sí FQ esta vacio entonces BQ esta vacio
2. Sí FQ tiene mas de un elemento entonces BQ esta vacio
*/
type Queue[T any] struct {
	fq, bq []T
}

func emptyQ[T any]() Queue[T] {
	return Queue[T]{}
}
func (q Queue[T]) isEmptyQ() bool {
	return len(q.fq) == 0
}
func (q *Queue[T]) enqueue(element T) {
	if len(q.fq) > 0 {
		q.bq = append(q.bq, element)
	} else {
		q.fq = append(q.fq, element)
	}
}

func (q Queue[T]) firstQ() T {
	return q.fq[0]
}

func (q *Queue[T]) dequeue() {
	q.fq = q.fq[1:]

	if len(q.fq) == 0 {
		q.fq = q.bq
		q.bq = make([]T, 0)
	}
}
