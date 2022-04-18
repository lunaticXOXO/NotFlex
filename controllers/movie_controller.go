package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id := r.Form.Get("ID")
	title := r.Form.Get("Title")
	release_date := r.Form.Get("Date")
	synopsis := r.Form.Get("Synopsis")
	director := r.Form.Get("Director")
	genre := r.Form.Get("Genre")
	actor := r.Form.Get("Actor")

	var movie Movie
	var response MovieResponse

	row := db.QueryRow("SELECT id FROM movie WHERE title=?", title)
	err = row.Scan(&movie.Id)

	if err != nil {
		_, errQuery := db.Exec("INSERT INTO movie(id,title,releaseDate,genre,actor,director,synopsis) VALUES (?,?,?,?,?,?,?)",
			id,
			title,
			release_date,
			genre,
			actor,
			director,
			synopsis,
		)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Add Movie!"
		} else {
			response.Status = 400
			response.Message = "Failed to Add New Movie!"
		}

	} else {
		response.Status = 400
		response.Message = "Title Already Exist!"
	}

	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	idloc, _ := strconv.Atoi(r.Form.Get("ID"))
	titleloc := r.Form.Get("Title")
	releasedateloc := r.Form.Get("Release Date")
	genreloc := r.Form.Get("Genre")
	synopsisloc := r.Form.Get("Synopsis")
	filmdirectorloc := r.Form.Get("Director")
	actorloc := r.Form.Get("Actor")

	_, errQuery := db.Exec("UPDATE movie SET title = ?, releaseDate = ?, genre = ?, actor = ?, director = ?, synopsis = ? WHERE id=?",
		titleloc,
		releasedateloc,
		genreloc,
		actorloc,
		filmdirectorloc,
		synopsisloc,
		idloc,
	)

	var response MovieResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Update Movie"
	} else {

		response.Status = 400
		response.Message = "Update Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ShowMovieByTitle(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	gettitle := r.URL.Query()["Title"]
	query := "SELECT * FROM movie"

	if gettitle != nil {
		query += " WHERE title LIKE '%" + gettitle[0] + "%'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var movie Movie
	var movies []Movie

	for rows.Next() {
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.ReleaseDate, &movie.Genre, &movie.Actor, &movie.Director, &movie.Synopsis); err != nil {
			log.Print(err.Error())
		} else {
			movies = append(movies, movie)
		}
	}

	var response MovieResponse
	if err == nil {
		response.Status = 200
		response.Message = "List of Movies:"
		response.Data = movies
	} else {
		response.Status = 400
		response.Message = "Error"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ShowMovieById(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	getid := vars["id"]

	rows, err := db.Query("SELECT * FROM movie WHERE id = ?", getid)

	if err != nil {
		log.Fatal(err)
	}

	var movie Movie
	var movies []Movie

	for rows.Next() {
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.ReleaseDate, &movie.Genre, &movie.Actor, &movie.Director, &movie.Synopsis); err != nil {
			log.Print(err.Error())
		} else {
			movies = append(movies, movie)
		}

	}

	var response MovieResponse

	if err == nil {
		response.Status = 200
		response.Message = "List of Movies:"
		response.Data = movies
	} else {
		response.Status = 400
		response.Message = "Error"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func SearchMovie(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	gettitle := r.URL.Query()["Title"]
	getfilmdirector := r.URL.Query()["Director"]
	releasedateloc := r.URL.Query()["Release Date"]
	genreloc := r.URL.Query()["Genre"]
	actorname := r.URL.Query()["Actor"]
	sinopsisloc := r.URL.Query()["Synopsis"]

	query := "SELECT * FROM movie"

	if gettitle != nil {
		query += " WHERE title LIKE '%" + gettitle[0] + "%'"
	} else if getfilmdirector != nil {
		query += " WHERE director LIKE '%" + getfilmdirector[0] + "%'"
	} else if releasedateloc != nil {
		query += " WHERE releaseDate LIKE '%" + releasedateloc[0] + "%'"
	} else if genreloc != nil {
		query += " WHERE genre LIKE '%" + genreloc[0] + "%'"
	} else if actorname != nil {
		query += " WHERE actor LIKE '%" + actorname[0] + "%'"
	} else if sinopsisloc != nil {
		query += " WHERE synopsis LIKE '%" + sinopsisloc[0] + "%'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var movie Movie
	var movies []Movie

	for rows.Next() {
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.ReleaseDate, &movie.Genre, &movie.Actor, &movie.Director, &movie.Synopsis); err != nil {
			log.Print(err.Error())
		} else {
			movies = append(movies, movie)
		}
	}

	var response MovieResponse

	if err == nil {
		response.Status = 200
		response.Message = "List of Movies:"
		response.Data = movies
	} else {
		response.Status = 400
		response.Message = "Error Get Movies!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func WatchMovie(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	judul := r.URL.Query()["Title"]

	rows, err := db.Query("SELECT * FROM movie WHERE title = ?", judul[0])
	if err != nil {
		log.Print(err)
	}

	var movie Movie
	var movies []Movie

	for rows.Next() {
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.ReleaseDate, &movie.Genre, &movie.Actor, &movie.Director, &movie.Synopsis); err != nil {
			log.Print(err.Error())
		} else {
			movies = append(movies, movie)
		}

	}

	_, errQuery := db.Exec("INSERT INTO historymovie(user_id, movie_id) VALUES(?,?)",
		user.Id,
		movie.Id,
	)

	var response MovieResponse

	if errQuery == nil {
		response.Status = 200
		response.Message = "Enjoy The Movie!"
		response.Data = movies
	} else {
		response.Status = 400
		response.Message = "Error"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
