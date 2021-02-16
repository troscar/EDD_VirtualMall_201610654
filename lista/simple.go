package lista

import (
	"fmt"
)

type nodo struct {
	siguiente, anterior *nodo
	calificacion int
	nombre , descripcion, telefono string
}

type Lst struct {
	raiz, ultimo *nodo
	size int
}

func NewLst() *Lst {
	return &Lst{nil,nil,0}
}

func (m *Lst)Insert(cali int,nom string,desc string, tel string) int {
	nuevo := &nodo{nil,nil,cali,nom,desc,tel}
	if m.raiz==nil {
		m.raiz = nuevo
		m.ultimo = nuevo
	}else{
		m.ultimo.siguiente = nuevo
		nuevo.anterior = m.ultimo
		m.ultimo = nuevo
	}
	m.size++
	return 1
}

func (m *Lst)Print()  {
	aux := m.raiz
	contador := 0
	for(aux != nil){
		fmt.Println(aux.nombre)
		aux = aux.siguiente
		contador++;
	}
}