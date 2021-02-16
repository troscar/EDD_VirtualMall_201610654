package matriz

import "fmt"

type Product struct {
	nombre string
	codigo int
	descripcion string
}

type nodo struct {
	//estos atributos son especificos para la matriz
	x, y     int      //saber en que cabecera estoy
	producto *Product //tipo de objeto
	izquierdo, derecho , arriba , abajo *nodo //nodos con los que nos deplazamos dentro de la matriz
	//ESTOS ATRIBUTOS SON ESPECIFICOS PARA LA LISTA
	header int // tipo interno de a cabecera
	siguiente, anterior *nodo // nodos con los qe vamos a desplazarnos dentro de las listas
}

type matriz struct {
	ls_h, lst_v *Lista
}

func nodoMatriz (x int ,y int,producto *Product) *nodo {
	return &nodo{x,y,producto,nil,nil,nil,nil,nil,nil,nil}
}

type Lista struct {
	first, last *nodo
}

func nodoLista(header int) *nodo {
	return &nodo{nil,nil,nil,nil,nil,nil,nil,header,nil,nil}
}

func newLista() *Lista {
	return &Lista{nil,nil}
}
func newMatriz() *matriz {
	return &matriz{newLista(),newLista()}
}

func (n *nodo) headerX()  int {return n.x}

func (n *nodo) headerY()  int {return n.y}

func (n *nodo) tostring()  string {return "Noombre: "+n.producto.nombre +"\n Descripcion: "+n.producto.descripcion}

var lista = &Lista{nil,nil}

func (l *Lista) Ordenar(nuevo *nodo)  {
	aux := l.first
	for (aux != nil){
		if nuevo.header > aux.header{
			aux = aux.siguiente
		}else{
			if aux == l.first{
				nuevo .siguiente = aux
				aux.anterior = nuevo
				l.first = nuevo
			}else{
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente  = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.last.siguiente =nuevo
	nuevo.anterior = l.last
	l.last = nuevo
}
func (l *Lista) Insert(header int)  {
	nuevo := nodoLista(header)
	if l.first == nil{
		l.first = nuevo
		l.last = nuevo
	}else{
		l.Ordenar(nuevo)
	}

}

func (l * Lista) Busqueda(header int) *nodo  {
	temp := l.first
	for temp != nil{
		if temp.header == header{
			return temp
		}
		temp = temp.siguiente
	}
	return  nil
}

func (l *Lista) Print()  {
	temp := l.first
	for temp != nil{
		fmt.Println("Cabecera: ",temp.header)
		temp = temp.siguiente
	}
}

func (m *matriz) Insert(product *Product,x int,y int)  {
	h := m.ls_h.Busqueda(x)
	v:= m.lst_v.Busqueda(y)

	if h==nil && v==nil{
		m.NoExisten(product,x,y)
	}else if h==nil && v !=nil{
		m.ExisteVertical(product,x,y)
	}else if h!=nil && v==nil{
		m.ExisteHorizontal(product,x,y)
	}else if h!= nil && v!= nil{
		m.Existen(product,x,y)
	}
}

func (m *matriz) NoExisten(product *Product, x int, y int) {
	m.ls_h.Insert(x) //insertamos en la lista que emula lacabecera horizontal
	m.lst_v.Insert(y) //insertamos en la lista que emula lacabecera vertical

	h:= m.ls_h.Busqueda(x) // vamos a buscar el nodo que acabamos de insertar para poder enlazar
	v:= m.lst_v.Busqueda(y) // vamos a buscar el nodo que acabamos de insertar para poder enlazar

	nuevo := nodoMatriz(x,y,product) // creamos nuevo nodo tipo matriz

	h.abajo = nuevo	//enlazamos el nodo horizontal hacia abajo
	nuevo.arriba = h // enlazamos el nodo nuevo hacia arriba

	v.derecho = nuevo //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierdo = v // enlazamos el nuevo nodo hacia la izquierda
}

func (m *matriz) ExisteVertical(product *Product, x int, y int) {
	
}

func (m *matriz) ExisteHorizontal(product *Product, x int, y int) {
	
}

func (m *matriz) Existen(product *Product, x int, y int) {
	
}
