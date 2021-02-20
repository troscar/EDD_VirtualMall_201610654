package matriz

import (
	"../listlist"
	"fmt"
	"os"
)

type NodoM struct {
	ll * listadelistas.LL
	siguiente, anterior, arriba, abajo *NodoM
	letra, categoria string
}

type Matriz struct {
	raiz *NodoM

}

func NewNodom(l, cate string) *NodoM {
	return &NodoM{/*listadelistas.NewLL()*/iniciar_lista(), nil, nil,nil,nil,l,cate}
}

func iniciar_lista() *listadelistas.LL  {
	ll := listadelistas.NewLL()
	ll.Insert(1,1,"","","")
	ll.Insert(2,2,"","","")
	ll.Insert(3,3,"","","")
	ll.Insert(4,4,"","","")
	ll.Insert(5,5,"","","")
	return ll

}

func NewMatriz() *Matriz {
	return &Matriz{NewNodom("raiz","raiz")}
}

func (m *Matriz) Buscar_letra(l string) *NodoM{
	aux := m.raiz.siguiente
	for aux!=nil{
		if aux.letra == l {
			return aux
		}
		aux = aux.siguiente
	}
	return nil
}

func (m *Matriz) Insertar_letra(l string) {
	aux := m.raiz.siguiente
	aux2:= m.raiz
	nuevo := NewNodom(l,"")
	for aux!=nil{
		if aux.letra > l{
			nuevo.siguiente = aux
			aux2.siguiente = nuevo
			aux.anterior  = nuevo
			nuevo.anterior = aux2
			break
		}
		aux = aux.siguiente
		aux2 = aux2.siguiente
	}
	if aux == nil{
		nuevo.anterior=aux2
		aux2.siguiente = nuevo
	}
}

func (m *Matriz) Buscar_categoria(c string) *NodoM{
	aux := m.raiz
	aux = aux.abajo
	for aux!=nil{
		if aux.categoria == c{
			return aux
		}
		aux = aux.abajo
	}
	return nil
}

func (m *Matriz) Insertar_categoria(cat string) {
	aux := m.raiz.abajo
	aux2:= m.raiz
	nuevo := NewNodom("",cat)
	for aux!=nil{
		if aux.categoria > cat{
			nuevo.abajo = aux
			aux2.abajo = nuevo
			aux.arriba  = nuevo
			nuevo.arriba = aux2
			break
		}
		aux = aux.abajo
		aux2 = aux2.abajo
	}
	if aux == nil{
		nuevo.arriba=aux2
		aux2.abajo = nuevo
	}
}

func (m *Matriz) Insert(let string, ria string){
	nuevoC := m.Buscar_categoria(ria)
	nuevoL := m.Buscar_letra(let)
	nuevo := NewNodom(let,ria)
	//fmt.Println(m.Buscar_letra(let))
	//fmt.Println(m.Buscar_categoria(ria))
	if nuevoL==nil && nuevoC==nil {
		//fmt.Println("entra")
		m.Insertar_categoria(ria)
		m.Insertar_letra(let)
		nuevoC = m.Buscar_categoria(ria)
		nuevoL = m.Buscar_letra(let)
		nuevo.arriba = nuevoL
		nuevo.anterior = nuevoC
		nuevoL.abajo = nuevo
		nuevoC.siguiente = nuevo
	}else if nuevoL!=nil && nuevoC==nil{
		m.Insertar_categoria(ria)
		m.Insertar_nodo_letra(let,ria,nuevo)
	}else if nuevoL==nil && nuevoC!=nil{
		m.Insertar_letra(let)
		m.Insertar_nodo_categoria(let,ria,nuevo)
	}else if nuevoL!=nil && nuevoC!=nil{
		existe := m.Buscar_Nodo(let, ria)
		if existe== nil {
			//fmt.Println("no existe")
			m.Insertar_letra_categoria(let,ria,nuevo)
		}else{
			//fmt.Println("ya existe")
		}
	}
}

func (m *Matriz) Buscar_Nodo(let string, ria string) *NodoM {
	aux1 := m.raiz
	aux2 := m.raiz
	for aux2!= nil {
		//fmt.Print( aux1.letra, "/",aux1.categoria,"-----------")
		if aux1.letra == let && aux1.categoria==ria{
			return aux1
		}
		if aux1.siguiente!=nil {
			aux1 = aux1.siguiente
		}else{
			aux2 = aux2.abajo
			aux1 = aux2
			//fmt.Println(" ")
		}
	}
	return nil
}

func (m *Matriz) Insertar_nodo_letra(let string,cate string,nuevo *NodoM)  {
	catego := m.Buscar_categoria(cate)
	aux := m.Buscar_letra(let)
	aux2 := aux
	aux = aux.abajo
	for aux!=nil{
		if aux.categoria > cate{
			nuevo.abajo = aux
			aux2.abajo = nuevo
			aux.arriba  = nuevo
			nuevo.arriba = aux2
			catego.siguiente = nuevo
			nuevo.anterior = catego
			break
		}
		aux = aux.abajo
		aux2 = aux2.abajo
	}
	if aux == nil{
		nuevo.arriba=aux2
		aux2.abajo = nuevo
		catego.siguiente = nuevo
		nuevo.anterior = catego

	}
}

func (m *Matriz) Insertar_nodo_categoria(let string, ria string, nuevo *NodoM) {
	letra := m.Buscar_letra(let)
	aux := m.Buscar_categoria(ria)
	aux2:= aux
	aux = aux.siguiente
	for aux!= nil{
		if aux.letra > let{
			nuevo.siguiente = aux
			aux2.siguiente = nuevo
			aux.anterior = nuevo
			nuevo.anterior = aux2
			letra.abajo = nuevo
			nuevo.arriba = letra
			break
		}
		aux = aux.siguiente
		aux2 = aux2.siguiente
	}
	if aux == nil {
		nuevo.anterior = aux2
		aux2.siguiente = nuevo
		letra.abajo = nuevo
		nuevo.arriba = letra
	}
}

func (m *Matriz) Insertar_letra_categoria(let string, cate string, nuevo *NodoM) {
	auxl := m.Buscar_letra(let)
	auxl2 := auxl
	auxl = auxl.abajo
	for auxl !=nil{
		if auxl.categoria > cate{
			break
		}
		auxl = auxl.abajo
		auxl2 = auxl2.abajo
	}
	if auxl != nil{
		nuevo.abajo = auxl
		auxl2.abajo = nuevo
		auxl.arriba  = nuevo
		nuevo.arriba = auxl2
	}else {
		auxl2.abajo = nuevo
		nuevo.arriba = auxl2
	}


	auxc := m.Buscar_categoria(cate)
	auxc2 := auxc
	auxc = auxc.siguiente
	for auxc != nil{
		if auxc.letra > let{
			break
		}
		auxc = auxc.siguiente
		auxc2 = auxc2.siguiente
	}
	if auxl != nil{
		nuevo.siguiente = auxc
		auxc2.siguiente = nuevo
		auxc.anterior = nuevo
		nuevo.anterior = auxc2
	}else {
		auxc2.siguiente = nuevo
		nuevo.anterior = auxc2
	}
}

func (m *Matriz) Insertar_listas (letra string , categoria string ,calificacion int,nom string,desc string, tel string)  {
	existe := m.Buscar_Nodo(letra, categoria)
	if existe != nil{
		existe.ll.Insert(calificacion,calificacion,nom,desc,tel)
	}
}

func (m *Matriz) Devuelve_LL (letra string , categoria string )  *listadelistas.LL  {
	existe := m.Buscar_Nodo(letra, categoria)
	if existe != nil{
		return m.Buscar_Nodo(letra,categoria).ll
	}
	return nil
}

func (m *Matriz) Print(){
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
	fmt.Println("----------------------------------------------------------------------")
	aux1 = m.raiz
	aux2 = m.raiz
	for aux2!= nil {
		fmt.Print( aux1.letra, "/",aux1.categoria,"-----------")
		if aux1.abajo!=nil {
			aux1 = aux1.abajo
		}else{
			aux2 = aux2.siguiente
			aux1 = aux2
			fmt.Println(" ")
		}
	}
}

func (m *Matriz) DotGraphviz(){
	f, err :=os.Create("Matrix.dot")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Fprintln(f,"digraph G {\n    node[style=\"filled\",shape= \"box\"]\n    graph[splines = \"ortho\"]\n}")
	f.Close()
}

func (m * Matriz) Arregllo()  {

}