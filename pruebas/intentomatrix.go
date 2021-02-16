package pruebas

import "fmt"
import "../listlist"

type NodoM struct {
	ll * listadelistas.LL
	siguiente, anterior, arriba, abajo *NodoM
	letra, categoria string
}
func NewNodom(l, cate string) *NodoM {
	return &NodoM{listadelistas.NewLL(), nil, nil,nil,nil,l,cate}
}
type Matrix struct {
	raiz *NodoM

}

func NewMatriz() *Matrix {
	return &Matrix{NewNodom("raiz","raiz")}
}

func (m *Matrix)Insert(let string, ria string){
	nuevoL := NewNodom(let,"")
	nuevoC := NewNodom("", ria)
	nuevo := NewNodom(let,ria)
	//temporal := m.raiz
	//INSERTA SI LA MATRIZ ESTA VACIA
	if m.raiz.siguiente==nil && m.raiz.abajo==nil {
		m.raiz.siguiente = nuevoL
		nuevoL.anterior = m.raiz
		m.raiz.abajo = nuevoC
		nuevoC.arriba = m.raiz
		nuevoL.abajo = nuevo
		nuevoC.siguiente = nuevo
	}
}
func (m *Matrix) Print(){
	aux1 := m.raiz
	aux2 := m.raiz
	for aux2!= nil {
		fmt.Print( aux1.letra, "/",aux1.categoria,"-----------")
		if aux1.siguiente!=nil {
			aux1 = aux1.siguiente
		}else{
			aux2 = aux2.abajo
			aux1 = aux2
			fmt.Println(" ")
		}
	}
}