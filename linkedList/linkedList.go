package linkedlist

/*
Dado un LinkedList con los campos head, size y end:
1. Sí el campo head != nil entonces end != nil
2. El campo 'size' no puede ser menor a 0
3. 'size' es igual a la cantidad de elementos del LinkedList
4. Sí 'size' es > 1 entonces 'head' y 'end' no pueden ser iguales
*/
type node[T comparable] struct {
	next    *node[T]
	element T
}

type LinkedList[T comparable] struct {
	head, tail *node[T]
	size       int
}

func emptyLL[T comparable]() LinkedList[T] {
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

func (l LinkedList[T]) containsLL(element T) bool {
	actual := l.head

	for actual != nil && actual.element != element {
		actual = actual.next
	}

	return actual != nil && actual.element == element
}

func (l LinkedList[T]) toList() []T {
	var list []T
	actual := l.head
	for actual != nil {
		list = append(list, actual.element)
		actual = actual.next
	}
	return list
}

// append element at index 'index'
// precond: index can't be greater than 'sizes' linkedlist
func (l *LinkedList[T]) insertAt(index int, element T) {
	if index > l.size {
		panic("index can't be greater than 'sizes'")
	}
	actual := l.head
	for i := 1; i < index; i++ {
		actual = actual.next
	}
	newNode := &node[T]{element: element, next: actual.next}
	actual.next = newNode
}

// Deletes the first element of the LinkedList
// If it is empty then it does nothing
func (l *LinkedList[T]) popLL() {
	if l.head != nil {
		l.head = l.head.next
	}
}
