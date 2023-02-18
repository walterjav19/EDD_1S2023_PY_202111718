package estructuras

import (
	"fmt"
	"roles"
)

type Node struct {
	data interface{}
	next *Node
}

type Queu struct {
	head *Node
	tail *Node
	size int
}

func (q *Queu) Enqueu(data interface{}) {
	newNode := &Node{data, nil} //direccion en memoria
	if q.tail == nil {
		q.head = newNode // el primero y el ultimo apuntan al nuevo nodo
		q.tail = newNode
	} else { //encolamos y el final el next del que era ultimo apunta al nuevo ultimo
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++ //en cualquier cosa se encolo +1
}

func (q *Queu) Dequeu() *roles.Student {
	if q.head == nil {
		return nil
	}
	data, ok := q.head.data.(*roles.Student)
	if !ok {
		return nil
	}
	q.head = q.head.next
	q.size--

	if q.head == nil {
		q.tail = nil
	}

	return data
}

func (q *Queu) Recorrer() {
	aux := q.head
	for aux != nil {
		fmt.Print(aux.data)
		fmt.Print("->")
		aux = aux.next
	}
}

func (q *Queu) Size() int {
	return q.size
}

func (q *Queu) IsEmpty() bool {
	return q.size == 0
}

func (q *Queu) Cabeza() *roles.Student {
	if s, ok := q.head.data.(*roles.Student); ok {
		return s
	}
	return nil
}
