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


type nodoLL struct {
	lst *lista.Lst
	siguiente, anterior *nodoLL
	index int
}

func newNodo(index int) *nodoLL {
	return &nodoLL{lista.NewLst(), nil, nil, index}
}

type LL struct{
	raiz, ultimo *nodoLL
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

func (ll *LL) Devolver_nodo_llista(index int)  *lista.Lst{
	aux := ll.raiz
	for aux!= nil {
		//fmt.Println("--------lista:", aux.index, "---------")
		if aux.index == index{
			return aux.lst
		}
		aux = aux.siguiente
	}
	return nil
}
