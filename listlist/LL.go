package listadelistas

import (
	"../lista"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
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


func (ll *LL) Insert(index int,cali int,nom string,desc string, tel string,let string,cate string)  {
	nuevo := newNodo(index)
	if ll.raiz == nil{
		ll.raiz = nuevo
		ll.ultimo = nuevo
		nuevo.lst.Insert(cali ,nom ,desc , tel ,let,cate)
	}else{
		aux := ll.raiz
		for(aux!=nil){
			if aux.index==index{
				aux.lst.Insert(cali ,nom ,desc , tel ,let,cate)
				return
			}
			aux = aux.siguiente
		}
		nuevo.lst.Insert(cali ,nom ,desc , tel ,let,cate)
		ll.ultimo.siguiente = nuevo
		nuevo.anterior= ll.ultimo
		ll.ultimo = nuevo
	}
	ll.size++
	return
}

func (ll *LL) InsertNodoLL(nodo *nodoLL,index int)  {
	nuevo := &nodoLL{nodo.lst,nil,nil,index}
	if ll.raiz == nil{
		ll.raiz = nuevo
		ll.ultimo = nuevo
	}else{
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
		fmt.Println("--------lLista:", aux.index, "---------")
		aux.lst.Print()
		aux = aux.siguiente
	}
}

func (ll *LL) DotGraphviz() {
	//ESCRIBIMOS EL ARCHIVO DOT
	f, err := os.Create("C:\\Users\\tracs\\Desktop\\Go\\PROYECTO\\graphviz\\LL.dot")
	if err != nil {
		fmt.Println(err)
	}
	cuerpo := "digraph G {\n    node[style=\"filled\",shape= \"record\"]\n    graph[splines = \"ortho\"]\n"

	////AQUI VAN LOS NODOS DE LA MATRIZ

	aux := ll.raiz
	conta := 0
	cuerpo = cuerpo + "\n node1[label=\""
	for aux != nil {
		//fmt.Println("--------lLista:", aux.index, "---------")
		cuerpo = cuerpo + " <f" + strconv.Itoa(conta) + "> " + strconv.Itoa(aux.index)

		//aux.lst.Print()
		aux = aux.siguiente
		conta++
		if aux != nil {
			cuerpo = cuerpo + " | "
		}
	}
	cuerpo = cuerpo + "\"]\n"

	aux = ll.raiz
	for aux != nil {
		cuerpo = cuerpo + aux.lst.Graph(aux.index)
		aux=aux.siguiente
	}

	///TERMINA EL ARCHIVO DOT
	cuerpo = cuerpo + "}"
	fmt.Fprintln(f,cuerpo)
	f.Close()

	//PASAR DE .DOT A PNG
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "C:\\Users\\tracs\\Desktop\\Go\\PROYECTO\\graphviz\\LL.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("C:\\Users\\tracs\\Desktop\\Go\\PROYECTO\\graphviz\\LL.png", cmd, os.FileMode(mode))
}

func (ll *LL) Devolver_nodo_llista(index int)  *nodoLL{
	aux := ll.raiz
	for aux!= nil {
		//fmt.Println("--------lista:", aux.index, "---------")
		if aux.index == index{
			return aux
		}
		aux = aux.siguiente
	}
	return nil
}

func (ll *LL) Devolver_lista(index int)  *lista.Lst{
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
