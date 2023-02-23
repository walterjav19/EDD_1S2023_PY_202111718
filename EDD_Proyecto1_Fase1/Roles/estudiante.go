package roles

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	top *stackNode // cima de la pila
	size int       // tamaño de la pila
}

type stackNode struct {
	value interface{} // valor del nodo
	prev  *stackNode  // nodo anterior
}



type Student struct {
	Nombre   string
	Apellido string
	Carnet   int
	Password string
	Pila *Stack
}

func (s *Student) ToString() string {
	return "Nombre: " + s.Nombre + " " + s.Apellido + ", Carnet: " + strconv.Itoa(s.Carnet) + "\n" +
		"***********************************************"
}

func (s *Student) Push(value interface{}) {
	newNode := &stackNode{value, s.Pila.top}
	s.Pila.top = newNode
	s.Pila.size++
}

func (s *Student) Pop() interface{} {
	if s.Pila.top == nil {
		return nil
	}

	value := s.Pila.top.value
	s.Pila.top = s.Pila.top.prev
	s.Pila.size--
	return value
}

func (s *Student) Recorrer() {
	if s.Pila.top == nil {
		return
	}

	aux := s.Pila.top
	for aux != nil {
		// mostrar el valor del nodo
		fmt.Println(aux.value)
		aux = aux.prev
	}
}

func (s *Student) DotPilaEstudiante(num int) string{
	if s.Pila.top == nil {
		return ""
	}
	text:=""
	enmedio:=""
	aux := s.Pila.top
	i:=0
	j:= num
	text+=fmt.Sprintf(`a%d->n%d
	`,s.Carnet,j)
	enmedio+=fmt.Sprintf("a%d,",s.Carnet)
	for aux != nil {
		text+=fmt.Sprintf(`n%d[label="%s"]
		`,j,aux.value)
		enmedio+=fmt.Sprintf("n%d,",j)
		if i!=s.Pila.size-1{
			text+=fmt.Sprintf(`n%d->n%d
			`,j,j+1)

		}
		j++
		i++
		aux = aux.prev
	}
	enmedio=strings.TrimRight(enmedio,",")
	rank:=fmt.Sprintf("{rank=same;%s}\n",enmedio)
	text+=rank
	return text
}
