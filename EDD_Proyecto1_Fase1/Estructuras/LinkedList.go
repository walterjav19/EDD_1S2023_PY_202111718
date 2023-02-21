package estructuras

import (
	"fmt"
	"menus"
	"roles"
	"os"
	"os/exec"
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

func (l *LinkedList) GenerarDotLista() {
	file, err := os.Create("C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Grafics/Lista.dot")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(`digraph LinkedList {
	rankdir=LR;
	node [shape=record];
	nullI [shape=none, label="null", style=bold, height=0, width=0];
	nullF [shape=none, label="null", style=bold, height=0, width=0];
	`)
	if err != nil {
		panic(err)
	}

	aux:=l.head
	for aux!=nil{
		_,err=file.WriteString(fmt.Sprintf(`a%d[label="%d\n%s %s"]
	`,aux.data.Carnet,aux.data.Carnet,aux.data.Nombre,aux.data.Apellido))
		if err != nil {
			panic(err)
		}

		if aux.next!=nil{
			_,err=file.WriteString(fmt.Sprintf(`a%d->a%d
	`,aux.data.Carnet,aux.next.data.Carnet))
			if err != nil {
				panic(err)
			}
		}
		
		if aux.prev!=nil{
			_,err=file.WriteString(fmt.Sprintf(`a%d->a%d
	`,aux.data.Carnet,aux.prev.data.Carnet))
			if err != nil {
				panic(err)
			}
		}

		if aux.next==nil{
			_,err=file.WriteString(fmt.Sprintf(`a%d->nullF
			`,aux.data.Carnet))
			if err != nil {
				panic(err)
			}
		}

		if aux.prev==nil{
			_,err=file.WriteString(fmt.Sprintf(`nullI->a%d[dir=back]
			`,aux.data.Carnet))
			if err != nil {
				panic(err)
			}
		}


		aux=aux.next
	}


	

	

	_, err = file.WriteString("}\n")
	if err != nil {
		panic(err)
	}

	RutaOrigen:="C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Grafics/Lista.dot"
	RutaDestino:="C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Reportes/Lista.png"
	//ejecutamos el comando para generar mi imagen ruta del dot 	
	cmd := exec.Command("dot", "-Tpng", RutaOrigen, "-o",RutaDestino )
	ex := cmd.Run()
	if ex != nil {
		fmt.Println(ex)
		os.Exit(1)
	}

}

func (l *LinkedList) Generarjson() {
	file, err := os.Create("C:/Users/USUARIO/Desktop/1Semestre_2023/EDD/LAB/EDD_1S2023_PY_202111718/EDD_Proyecto1_Fase1/Reportes/Alumnos.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(`{
	"alumnos": [
		`)
	if err != nil {
		panic(err)
	}

	aux:=l.head
	for aux!=nil{
		if aux.next!=nil{
			_,err=file.WriteString(fmt.Sprintf(`{
			"nombre": "%s %s",
			"carnet": %d,
			"password": "%s",
			"Carpeta_Raiz": "/"
		},
	    `,aux.data.Nombre,aux.data.Apellido,aux.data.Carnet,aux.data.Password))
		if err!=nil{
			panic(err)
		}

		
	}else{
		_,err=file.WriteString(fmt.Sprintf(`{
			"nombre": "%s %s",
			"carnet": %d,
			"password": "%s",
			"Carpeta_Raiz": "/"
		}
	`,aux.data.Nombre,aux.data.Apellido,aux.data.Carnet,aux.data.Password))
		if err!=nil{
			panic(err)
		}		
	}
	aux=aux.next
	}

	

	

	_, err = file.WriteString("]\n}")
	if err != nil {
		panic(err)
	}

}