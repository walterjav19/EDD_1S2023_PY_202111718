package estructuras
import "fmt"

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