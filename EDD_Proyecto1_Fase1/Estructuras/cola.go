package estructuras

import (
	"fmt"
	"roles"
	"os"
	"os/exec"
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

func (q *Queu) GenerarDotCola() {
	file, err := os.Create("C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Grafics/cola.dot")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("digraph G {\n"+"rankdir=LR;\n"+"node [shape=record];\n")
	if err != nil {
		panic(err)
	}

	aux := q.head
	for aux != nil {
		if s, ok := aux.data.(*roles.Student); ok {
			_, err = file.WriteString(fmt.Sprintf("a%d [label=<<TABLE border=\"0\"><TR><TD ALIGN=\"CENTER\">%d</TD></TR><TR><TD ALIGN=\"CENTER\">%s %s</TD></TR></TABLE>>];\n", s.Carnet, s.Carnet, s.Nombre, s.Apellido))
			if err != nil {
				panic(err)
			}

			if aux.next != nil {
				if nextS, ok := aux.next.data.(*roles.Student); ok {
					_, err = file.WriteString(fmt.Sprintf("a%d -> a%d;\n", s.Carnet, nextS.Carnet))
					if err != nil {
						panic(err)
					}
				}
			}
		}
		aux = aux.next
	}

	_, err = file.WriteString("}\n")
	if err != nil {
		panic(err)
	}

	RutaOrigen:="C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Grafics/cola.dot"
	RutaDestino:="C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Reportes/cola.png"
	//ejecutamos el comando para generar mi imagen ruta del dot 	
	cmd := exec.Command("dot", "-Tpng", RutaOrigen, "-o",RutaDestino )
	ex := cmd.Run()
	if ex != nil {
		fmt.Println(ex)
		os.Exit(1)
	}


}