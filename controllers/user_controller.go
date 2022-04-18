package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

var user User

//Register
func RegistrationUser(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id := r.Form.Get("ID")
	fullnameloc := r.Form.Get("Name")
	emailloc := r.Form.Get("Email")
	passloc := r.Form.Get("Password")
	dob := r.Form.Get("Date of Birth")
	countrylog := r.Form.Get("Country")
	genderlog := r.Form.Get("Gender")

	var response UserResponse
	var user User

	row := db.QueryRow("SELECT id FROM users WHERE email=?", emailloc)
	err = row.Scan(&user.Id)

	if err != nil {
		_, errQuery := db.Exec("INSERT INTO users(id,fullName,email,password,dateofbirth,country,gender,user_type) VALUES (?,?,?,?,?,?,?,'M')",
			id,
			fullnameloc,
			emailloc,
			passloc,
			dob,
			countrylog,
			genderlog,
		)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Registration Success!"
		} else {
			response.Status = 400
			response.Message = "Failed to Register!"
		}

	} else {
		response.Status = 400
		response.Message = "Email Already Registered!"
	}

	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Login
func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	emailLog := r.URL.Query()["Email"]
	passLog := r.URL.Query()["Password"]

	row := db.QueryRow("SELECT * FROM users WHERE email=? AND password=?", emailLog[0], passLog[0])

	if err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Dateofbirth, &user.Country, &user.Gender, &user.User_type); err != nil {
		sendErrorResponse(w)
	} else {
		generateToken(w, user.Id, user.FullName, user.User_type)
		sendSuccessResponse(w)
	}
}

//Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)
	sendSuccessResponse(w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	idUser := user.Id
	fullnameloc := r.Form.Get("Name")
	tgllahirloc := r.Form.Get("Date of Birth")
	genderloc := r.Form.Get("Gender")

	_, errQuery := db.Exec("UPDATE users SET fullName = ?, dateofbirth = ?, gender = ? WHERE id = ?",
		fullnameloc,
		tgllahirloc,
		genderloc,
		idUser,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Update New Profile!"
	} else {
		response.Status = 400
		response.Message = "Failed to Update!"
	}

	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	user_id := user.Id

	rows, err := db.Query("SELECT historymovie.movie_id,movie.title,movie.director FROM historyMovie JOIN movie ON movie.id = historymovie.movie_id WHERE historymovie.user_id = ?", user_id)
	if err != nil {
		log.Print(err)
	}

	var history MovieHistory
	var histories []MovieHistory

	for rows.Next() {
		if err := rows.Scan(&history.Id, &history.Title, &history.Director); err != nil {
			log.Print(err.Error())
		} else {
			histories = append(histories, history)
		}
	}

	var response MovieHistoryResponse

	if err == nil {
		response.Status = 200
		response.Message = "Watch History:"
		response.Data = histories
	} else {
		response.Status = 400
		response.Message = "Error"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response UserResponse
	response.Status = 400
	response.Message = "Unauthorized Access"
	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response UserResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response UserResponse
	response.Status = 400
	response.Message = "Error"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
