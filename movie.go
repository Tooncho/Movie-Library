//Este sera el modelo de la pelicula

package main

//Definimos el struct
type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

//La estructura Movies sera un conjunto de Movie
type Movies []Movie
