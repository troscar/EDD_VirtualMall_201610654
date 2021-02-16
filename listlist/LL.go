package listadelistas

import (
	"../lista"
	"fmt"
)
//prueba en main
//ll := listadelistas.NewLL()
//ll.Insert(1,"")
//ll.Insert(2,"")
//ll.Insert(3,"")
//ll.Insert(4,"")
//ll.Insert(5,"")
//ll.Print()


type nodo struct {
	lst *lista.Lst
	siguiente, anterior *nodo
	index int
}

func newNodo(index int) *nodo {
	return &nodo{lista.NewLst(), nil, nil, index}
}

type LL struct{
	raiz, ultimo *nodo
	size int
}

func NewLL() *LL {
	return &LL{raiz: nil, ultimo: nil, size: 0}
}


func (ll *LL) Insert(index int,cali int,nom string,desc string, tel string)  {
	nuevo := newNodo(index)
	if ll.raiz == nil{
		ll.raiz = nuevo
		ll.ultimo = nuevo
		nuevo.lst.Insert(cali ,nom ,desc , tel )
	}else{
		aux := ll.raiz
		for(aux!=nil){
			if aux.index==index{
				aux.lst.Insert(cali ,nom ,desc , tel )
				return
			}
			aux = aux.siguiente
		}
		nuevo.lst.Insert(cali ,nom ,desc , tel )
		ll.ultimo.siguiente = nuevo
		nuevo.anterior= ll.ultimo
		ll.ultimo = nuevo
	}
	ll.size++
	return
}

func (ll *LL) Print(){
	aux := ll.raiz
	for aux!= nil {
		fmt.Println("--------lista:", aux.index, "---------")
		aux.lst.Print()
		aux = aux.siguiente
	}
}