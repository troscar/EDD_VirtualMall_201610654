package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	request()
}
//todo el documento JSON
type Sobre struct {
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

func homepage(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Welcome Servidor go")
}

func getArreglo(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"[1,2,3,4]")
}
func settodopost(w http.ResponseWriter,r *http.Request)  {
	body, _ := ioutil.ReadAll(r.Body)
	var request Sobre
	json.Unmarshal(body,&request)
	fmt.Println(request)
	fmt.Fprintf(w,"OK")
}

func request(){
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/",homepage)
	myrouter.HandleFunc("/getArreglo",getArreglo).Methods("GET")
	myrouter.HandleFunc("/settodopost",settodopost).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",myrouter))

}