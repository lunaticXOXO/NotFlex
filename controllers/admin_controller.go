package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	email := r.URL.Query()["Email"]
	query := "SELECT id,fullName,email,password,dateofbirth,country,gender FROM users"

	if email != nil {
		query += " WHERE email ='" + email[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var user UserSearch
	var users []UserSearch

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Dateofbirth, &user.Country, &user.Gender); err != nil {
			log.Print(err.Error())
		} else {
			users = append(users, user)
		}
	}

	var response UserSearchResponse
	if err == nil {
		response.Status = 200
		response.Message = "Success Get User!"
		response.Data = users
	} else {
		response.Status = 400
		response.Message = "Error"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func SuspendUser(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	vars := mux.Vars(r)
	email := vars["email"]

	_, errQuery := db.Exec("DELETE m.* FROM membership m JOIN users u ON m.user_id = u.id WHERE u.email = ?",
		email,
	)

	var response SubscriptionResponse

	if errQuery == nil {
		response.Status = 200
		response.Message = "Suspend Success!"
	} else {
		response.Status = 400
		response.Message = "Suspend Failed!"
	}

	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
