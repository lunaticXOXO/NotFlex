package controllers

import "time"

type Movie struct {
	Id          int       `form : "id" json : "id"`
	Title       string    `form : "title" json : "title"`
	ReleaseDate time.Time `form : "releaseDate" json : "releaseDate"`
	Synopsis    string    `form : "synopsis" json : "synopsis"`
	Director    string    `form : "director" json : "director"`
	Genre       string    `form : "genre" json : "genre"`
	Actor       string    `form : "actor" json : "actor"`
}

type MovieResponse struct {
	Status  int     `form : "status" json : "status"`
	Message string  `form : "message" json : "message"`
	Data    []Movie `form : "data" json : "data"`
}
