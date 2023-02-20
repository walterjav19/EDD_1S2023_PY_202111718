package estructuras

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)


type NodeStack struct {
	data interface{}
	next *NodeStack
}

type Stack struct {
	top *NodeStack
}

func (p *Stack) Push(valor interface{}) {
	Caja := &NodeStack{valor, p.top} //se coloca caja y abajo apunta al antiguo tope
	p.top = Caja
}

func (p *Stack) Pop() interface{}{
	if p.top ==nil{//vacia
		return nil
	}
	caja:=p.top.data
	p.top=p.top.next
	return caja
}

func (p *Stack) Top() interface{}{
	if p.top ==nil{//vacia
		return nil
	}
	return p.top
}

func (p *Stack) Recorrer(){
	aux:=p.top
	for aux!=nil{
		fmt.Println(aux.data)
		aux=aux.next
	}
}

func (p *Stack) CrearDot() {
	f, err := os.Create("C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Grafics/PilaAdmin.dot")
	if err != nil {
		fmt.Println("Error al crear archivo: ", err)
		return
	}

	defer f.Close()

	_, err = f.WriteString("digraph Pila {\n")
	if err != nil {
		fmt.Println("Error al escribir en archivo: ", err)
		return
	}

	// Escribir la configuración de nodos y aristas
	_, err = f.WriteString("node [shape=box, width=1.2, height=0.6, style=filled, fillcolor=lightgray];\n"+"rankdir=LR;\n"+"nodesep=0;\n")
	if err != nil {
		fmt.Println("Error al escribir en archivo: ", err)
		return
	}

	// Crear una copia de la pila
	copyStack := &Stack{}
	aux := p.top
	for aux != nil {
		copyStack.Push(aux.data)
		aux = aux.next
	}

	// Recorrer la copia de la pila a la inversa y escribir los nodos en el archivo Dot
	aux = copyStack.top
	i := 0
	for aux != nil {
		//fecha en el formato que indica
		now := time.Now().Format("2006-01-02 15:04:05")

		// Escribir el nodo con la fecha y hora
		label := fmt.Sprintf("%v\n%s", aux.data, now)

		// Escribir el nodo
		_, err = f.WriteString(fmt.Sprintf("n%d [label=\"%s\"];\n", i, label))
		if err != nil {
			fmt.Println("Error al escribir en archivo: ", err)
			return
		}
		aux = aux.next
		i++
	}

	_, err = f.WriteString("}\n")
	if err != nil {
		fmt.Println("Error al escribir en archivo: ", err)
		return
	}


	RutaOrigen:="C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Grafics/PilaAdmin.dot"
	RutaDestino:="C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Reportes/PilaAdmin.png"
	//ejecutamos el comando para generar mi imagen ruta del dot 	
	cmd := exec.Command("dot", "-Tpng", RutaOrigen, "-o",RutaDestino )
	ex := cmd.Run()
	if ex != nil {
		fmt.Println(ex)
		os.Exit(1)
	}
}
