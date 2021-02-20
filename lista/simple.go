package lista

import (
	"fmt"
)

type nodoL struct {
	siguiente, anterior *nodoL
	calificacion int
	nombre , descripcion, telefono string
}

type Lst struct {
	raiz, ultimo *nodoL
	size int
}

func NewLst() *Lst {
	return &Lst{nil,nil,0}
}

func (m *Lst)Insert(cali int,nom string,desc string, tel string) int {
	nuevo := &nodoL{nil,nil,cali,nom,desc,tel}
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

func (m *Lst) Devuelve_nodo_lista(index int) *nodoL {
	aux := m.raiz
	for(aux != nil){
		return aux
		aux = aux.siguiente
	}
	return nil
}