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
