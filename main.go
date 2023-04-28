package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Creamos un Router y con las url amigables
	router := NewRouter()

	fmt.Println("El servidor esta corriendo en localhost:8080")

	//Para levantar el servidor le pasamos el objeto router que creamos
	server := http.ListenAndServe(":8080", router)

	//Si tiene un error
	log.Fatal(server)

}
