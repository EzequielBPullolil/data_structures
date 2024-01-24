package linkedlist

/*
Dado un LinkedList con los campos head, size y end:
1. Sí el campo head != nil entonces end != nil
2. El campo 'size' no puede ser menor a 0
3. 'size' es igual a la cantidad de elementos del LinkedList
4. Sí 'size' es > 1 entonces 'head' y 'end' no pueden ser iguales
*/
type node[T any] struct {
	next    *node[T]
	element T
}

type LinkedList[T any] struct {
	head, tail *node[T]
	size       int
}

func emptyLL[T any]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l LinkedList[T]) isEmptyLL() bool {
	return l.size == 0
}

func (l LinkedList[T]) sizeLL() int {
	return l.size
}

func (l *LinkedList[T]) appendLL(element T) {
	nodeToAppend := &node[T]{
		next:    nil,
		element: element,
	}
	if l.head == nil {
		l.head = nodeToAppend
		l.tail = nodeToAppend
	} else {
		nodeToAppend.next = l.head
		l.head = nodeToAppend
	}

	l.size += 1
}

func (l *LinkedList[T]) consLL(element T) {
	nodeToAppend := &node[T]{
		next:    nil,
		element: element,
	}

	if l.tail == nil {
		l.head = nodeToAppend
		l.tail = nodeToAppend
	} else {
		l.tail.next = nodeToAppend
		l.tail = nodeToAppend
	}

	l.size += 1
}


