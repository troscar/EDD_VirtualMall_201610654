package main

import (
	"./matriz"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//var L = lista.NewLst()
//var Ll = listadelistas.NewLL()
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
	arre := Mm.Arregllo()
	arre.DotGraphviz()
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
				//fmt.Println("SI HAY TIENDAS")
				for k := 0; k < len(Tiendas); k++  {
					Calificacion := Tiendas[k].Calificacion
					Nombre := Tiendas[k].Nombre
					Descripcion := Tiendas[k].Descripcion
					Contacto := Tiendas[k].Contacto
					//fmt.Println("Tienda: "+Nombre+", Calificacion"+strconv.Itoa(Calificacion)+". Contacto: "+Contacto+", Depa: "+departamento+", Letra: "+letra+"\n Descripcion: "+Descripcion+"\n")
					Mm.Insertar_listas(letra,departamento,Calificacion,Nombre,Descripcion,Contacto)
				}
			}else{
				//fmt.Println("NO HAY TIENDAS")
			}
		}
	}
	//Mm.Print()
	//arre := Mm.Arregllo()
	//arre.DotGraphviz()
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
	myrouter.HandleFunc("/id/0",getid).Methods("GET")
	myrouter.HandleFunc("/guardarjson",getguardar).Methods("GET")
	myrouter.HandleFunc("/cargartienda",cargartienda).Methods("POST")
	myrouter.HandleFunc("/TiendaEspecifica",TiendaEspecifica).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",myrouter))

}

func getguardar(writer http.ResponseWriter, h *http.Request) {
	crear_json, _ := json.Marshal(sobre)
	convertir_a_cadena := string(crear_json)
	fmt.Println(convertir_a_cadena)
	f, err :=os.Create("C:\\Users\\tracs\\Desktop\\Go\\PROYECTO\\graphviz\\practica.json")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Fprintln(f,convertir_a_cadena)
}

func getid(writer http.ResponseWriter, h *http.Request) {
	arre := Mm.Arregllo()
	listado := arre.Devolver_lista(5)
	lista := listado.Devuelve_arreglo(listado)
	crear_json, _ := json.Marshal(lista)
	convertir_a_cadena := string(crear_json)
	fmt.Println(lista)
	fmt.Fprintln(writer,convertir_a_cadena)
}