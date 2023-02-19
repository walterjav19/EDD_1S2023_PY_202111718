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

func (l *LinkedList) Validar(carne int,password string) bool{
	aux := l.head
	for aux != nil {
		if aux.data.Carnet==carne && password==aux.data.Password{
			return true
		}
		aux = aux.next
	}
	return false
}

func (l *LinkedList) ObtenerCarnet(carne int) bool{
	aux := l.head
	for aux != nil {
		if aux.data.Carnet==carne{
			return true
		}
		aux = aux.next
	}
	return false
}

func (l *LinkedList) ObtenerPassword(pass string) bool{
	aux := l.head
	for aux != nil {
		if aux.data.Password==pass{
			return true
		}
		aux = aux.next
	}
	return false
}

func (l *LinkedList) GetStudent(carne int,password string) *roles.Student {
    aux := l.head
    for aux != nil {
        if aux.data.Carnet == carne && password==aux.data.Password {
            return aux.data
        }
        aux = aux.next
    }
    return nil
}

