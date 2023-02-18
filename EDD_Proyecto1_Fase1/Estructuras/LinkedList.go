package estructuras

import (
	"fmt"
	"menus"
	"roles"
)

type node struct {
	data *roles.Student //Estudiante
	next *node          //apuntador siguiente
	prev *node          //apuntador anterior
}
type LinkedList struct {
	head *node //primero
	tail *node //ultimo
	size int   //tamaño
}

func (l *LinkedList) Append(data *roles.Student) {
	newNode := &node{data, nil, l.tail} //el next del ultimo siempre a nil
	if l.tail != nil {                  //si hay nodo
		l.tail.next = newNode //el ultimo apuntara al nuevo ultimo
	} else { //si no hay nodo
		l.head = newNode //la cabeza al primero
	}
	l.tail = newNode //el ultimo siempre sera el nodo a insertar
	l.size++         //aumentamos el tamaño
	/*Estado1 A->nil
	Estado2 A<->B->nil
	Estado3 A<->B<->C->nil*/
}

func (l *LinkedList) Empty() bool {
	return l.size == 0
}

func (l *LinkedList) Recorrer() {
	if l.Empty() {
		menus.Vacia()
		return
	}

	aux := l.head
	for aux != nil {
		fmt.Println(aux.data.ToString())
		aux = aux.next
	}
}

func (l *LinkedList) Sort() {
	if l.size < 2 { //no hace falta ordenar
		return
	}

	var min *node
	var aux *node

	for i := l.head; i.next != nil; i = i.next {
		min = i
		for j := i.next; j != nil; j = j.next {
			if min.data.Carnet > j.data.Carnet {
				min = j
			}
		}

		if min != i {
			// Intercambiar valores
			aux = &node{min.data, i.next, i.prev}
			min.data, i.data = i.data, min.data

			if i.prev != nil {
				i.prev.next = &node{i.data, aux.next, aux.prev}
			} else {
				l.head = &node{i.data, aux.next, aux.prev}
			}

			if min.next != nil {
				min.next.prev = &node{min.data, aux.next, aux.prev}
			} else {
				l.tail = &node{min.data, aux.next, aux.prev}
			}
		}
	}
}
