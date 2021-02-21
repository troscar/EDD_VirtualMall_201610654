package lista

import (
	"fmt"
	"strconv"
)

type nodoL struct {
	siguiente, anterior *nodoL
	calificacion , indice int
	nombre , descripcion, telefono ,letra, categoria string
}

type Lst struct {
	raiz, ultimo *nodoL
	size int
}

func NewLst() *Lst {
	return &Lst{nil,nil,0}
}

func (m *Lst)Insert(cali int,nom string,desc string, tel string, let string , cate string) int {
	nuevo := &nodoL{nil,nil,cali,m.size,nom,desc,tel,let,cate}
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
	for(aux != nil){
		fmt.Println(strconv.Itoa(aux.indice) +" "+aux.nombre)
		aux = aux.siguiente
	}
}
func (m *Lst)Graph(n int) string  {
	aux := m.raiz
	lista := ""
	contador := 0
	for(aux != nil){
		if contador == 0 {
			lista = lista + "nodo"+strconv.Itoa(n)+"l"+strconv.Itoa(aux.indice)+"[label=\"" + aux.nombre + "\"]\n"
			lista = lista + "node1:f"+strconv.Itoa(n) +" -> " + "nodo"+strconv.Itoa(n)+"l"+strconv.Itoa(aux.indice)+"\n"
			aux = aux.siguiente
			contador++
		}else{
			//fmt.Println(aux.nombre)
			lista = lista + "nodo"+strconv.Itoa(n)+"l"+strconv.Itoa(aux.indice)+"[label=\"" + aux.nombre + "\"]\n"
			lista = lista + "nodo"+strconv.Itoa(n)+"l"+strconv.Itoa(aux.indice)+"->" + "nodo"+strconv.Itoa(n)+"l"+strconv.Itoa(aux.anterior.indice)+"\n"
			aux = aux.siguiente
		}

	}
	return lista
}

type Tienda struct {
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
}

func (m *Lst) Devuelve_arreglo(l *Lst) []Tienda {
	aux := m.raiz
	cont := 0
	for(aux != nil){
		cont ++
		aux = aux.siguiente
	}
	var list []Tienda
	for(aux != nil){
		fmt.Println(strconv.Itoa(aux.indice) +" "+aux.nombre)
		nodito := Tienda{aux.nombre,aux.descripcion,aux.telefono,aux.calificacion}
		list = append(list,nodito)
		aux = aux.siguiente
	}
	return list
}