package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"./matriz"
	"./lista"
	"./listlist"
	"strconv"
)

var L = lista.NewLst()
var Ll = listadelistas.NewLL()
var Mm = matriz.NewMatriz()
var sobre archivo

func main() {
	request()
}


//todo el documento JSON
type archivo struct {
	Datos []Dato
}
//Nombre de la letra
type Dato struct {
	Indice string
	Departamentos []Departamento
}
//Nombre del departamento
type Departamento struct {
	Nombre string
	Tiendas []Tienda
}
//calificacion e informacion de la tienda
type Tienda struct {
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
}

type Búsqueda_de_posición_específica struct {
	Departamento string
	Nombre string
	Calificacion int
}

func homepage(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Welcome Servidor go")
}

func getArreglo(w http.ResponseWriter,r *http.Request)  {
	Mm.DotGraphviz()

}

func cargartienda(w http.ResponseWriter,r *http.Request)  {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body,&sobre)
	//fmt.Println(request)
	//json.NewEncoder(w).Encode(request)
	//fmt.Fprintf(w,"OK")
	Letras := sobre.Datos
	for i := 0; i < len(Letras); i++ {
		letra := Letras[i].Indice
		Departamentos := Letras[i].Departamentos
		for j := 0; j < len(Departamentos); j++  {
			departamento := Departamentos[j].Nombre
			Mm.Insert(letra,departamento)
			Tiendas := Departamentos[j].Tiendas
			if len(Tiendas)!= 0 {
				fmt.Println("SI HAY TIENDAS")
				for k := 0; k < len(Tiendas); k++  {
					Calificacion := Tiendas[k].Calificacion
					Nombre := Tiendas[k].Nombre
					Descripcion := Tiendas[k].Descripcion
					Contacto := Tiendas[k].Contacto
					fmt.Println("Tienda: "+Nombre+", Calificacion"+strconv.Itoa(Calificacion)+". Contacto: "+Contacto+", Depa: "+departamento+", Letra: "+letra+"\n Descripcion: "+Descripcion+"\n")

				}
			}else{
				fmt.Println("NO HAY TIENDAS")
			}


		}
	}
	Mm.Print()
	arre := Mm.Arregllo()
	arre.DotGraphviz()
}

func TiendaEspecifica(w http.ResponseWriter,r *http.Request)  {
	body, _ := ioutil.ReadAll(r.Body)
	var especifica Búsqueda_de_posición_específica
	json.Unmarshal(body,&especifica)
	//fmt.Println(tienda)
	//json.NewEncoder(w).Encode(tienda)
}

func request(){
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/",homepage)
	myrouter.HandleFunc("/getArreglo",getArreglo).Methods("GET")
	myrouter.HandleFunc("/cargartienda",cargartienda).Methods("POST")
	myrouter.HandleFunc("/TiendaEspecifica",TiendaEspecifica).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",myrouter))

}