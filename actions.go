//Aqui estaran los metodos

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Para conectarnos a MongoDB
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

var collection = getSession().DB("curso_go").C("movies")

var movies = Movies{
	Movie{"Sin limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo gas", 2005, "Juan Antonio"},
}

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

// Creamos la respuesta de la pagina requerida
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Que raro esto che")
}

func MovieList(w http.ResponseWriter, r *http.Request) {

	var results []Movie
	err := collection.Find(nil).All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}

	responseMovies(w, 200, results)

}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	//Para recoger los parametros por url
	params := mux.Vars(r)
	movie_id := params["id"]

	//Comprobamos si es un objeto hexa
	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, results)

}

func MovieAdd(w http.ResponseWriter, r *http.Request) {

	//Recogemos los datos del body y lo decodificamos con json
	decoder := json.NewDecoder(r.Body)

	//Variable para los datos que recogemos y los decodificamos
	var movie_data Movie
	err := decoder.Decode(&movie_data)

	//Comprobamos si hay error
	if err != nil {
		panic(err)
	}

	//Cerramos la lectura del body
	defer r.Body.Close()

	//Guardamos la informacion en la BDD
	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseMovie(w, 200, movie_data)

}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	//Para recoger los parametros por url
	params := mux.Vars(r)
	movie_id := params["id"]

	//Comprobamos si es un objeto hexa
	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	//Enviamos el objeto con datos a actualizar
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}
	err = collection.Update(document, change)

	if err != nil {
		panic(err)
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movie_data)

}

type Message struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func MovieRemove(w http.ResponseWriter, r *http.Request) {
	//Para recoger los parametros por url
	params := mux.Vars(r)
	movie_id := params["id"]

	//Comprobamos si es un objeto hexa
	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	results := Message{"success", "La pelicula con ID " + movie_id + " ha sido borrada correctamente"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
