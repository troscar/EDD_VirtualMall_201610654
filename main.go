package main

import "./matriz"
import "./listlist"

func main() {
	//l := lista.NewLst()
	ll := listadelistas.NewLL()
	mm := matriz.NewMatriz()
	mm.Insert("A","hogar")
	mm.Insert("C","jardin")
	mm.Insert("E","garage")
	mm.Insert("C","deporte")
	mm.Insert("F","deporte")
	mm.Insert("F","jardin")
	mm.Insert("C","garage")
	ll= mm.Arregllo()
	ll.DotGraphviz()
	mm.DotGraphviz()

	//mm.Insertar_listas("A","hogar",4,"Berna","licuados locos","22889318")
	//ll = mm.Devuelve_LL("A","hogar")
	//l = mm.Devuelve_LL("A","hogar").Devolver_nodo_llista(4)
	//l.Print()
	ll.Print()
	//mm.Print()


}