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

func (q *Queu) Dequeu() interface{} { //nos retornara la info del nodo
	if q.head == nil { //esta vacia
		return nil
	}
	data := q.head.data  //el nodo que se borra
	q.head = q.head.next //cabeza sera el siguiente en ser atendido
	q.size--             //si llego aca se disminuye

	if q.head == nil { //si borramos al ultimo entonces el final tambien debe ser nil
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

func (q *Queu) Cabeza() string {
	if student, ok := q.head.data.(roles.Student); ok {
		return student.Nombre + " " + student.Apellido
	}
	return "" // o algún valor por defecto si el tipo de dato no es Student o si es nulo
}
