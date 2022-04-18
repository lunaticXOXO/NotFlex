package main

import (
	"fmt"
	"log"
	"net/http"

	c "github.com/TubesPBP/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	//A -> Admin
	//M -> Member

	//Role Administrator
	router.HandleFunc("/searchUser", c.Authenticate(c.GetUserByEmail, "A")).Methods("GET")
	router.HandleFunc("/suspendUser/{email}", c.Authenticate(c.SuspendUser, "A")).Methods("DELETE")
	router.HandleFunc("/addNewMovie", c.Authenticate(c.AddMovie, "A")).Methods("POST")
	router.HandleFunc("/updateMovie", c.Authenticate(c.UpdateMovie, "A")).Methods("PUT")
	router.HandleFunc("/searchMovieAdmin", c.Authenticate(c.ShowMovieByTitle, "A")).Methods("GET")
	router.HandleFunc("/searchMovieAdmin/{id}", c.Authenticate(c.ShowMovieById, "A")).Methods("GET")

	//Role Member
	router.HandleFunc("/updateProfile", c.Authenticate(c.UpdateUser, "M")).Methods("PUT")
	router.HandleFunc("/searchMovieMember", c.Authenticate(c.SearchMovie, "M")).Methods("GET")
	router.HandleFunc("/addSubscription", c.Authenticate(c.AddSubscription, "M")).Methods("POST")
	router.HandleFunc("/deleteSubscription", c.Authenticate(c.DeleteSubscription, "M")).Methods("DELETE")
	router.HandleFunc("/watchMovie", c.Authenticate(c.WatchMovie, "M")).Methods("GET")
	router.HandleFunc("/getHistory", c.Authenticate(c.GetHistory, "M")).Methods("GET")

	//Login,Logout,Register
	router.HandleFunc("/register", c.RegistrationUser).Methods("POST")
	router.HandleFunc("/login", c.Login).Methods("GET")
	router.HandleFunc("/logout", c.Logout).Methods("GET")

	corseHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := corseHandler.Handler(router)

	http.Handle("/", handler)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
