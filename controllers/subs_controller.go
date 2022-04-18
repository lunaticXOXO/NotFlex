package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func AddSubscription(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id := r.Form.Get("ID")
	jenisloc := r.Form.Get("Subscription Plan")
	duration := time.Now().AddDate(0, 1, 0)
	price := 0
	noKartuKredit := r.Form.Get("CCNumber")
	masaBerlaku := r.Form.Get("Valid Thru")
	kodeCVC := r.Form.Get("CVC")

	if jenisloc == "Basic" {
		price = 50000
	} else if jenisloc == "Premium" {
		price = 100000
	} else {
		sendErrorResponse(w)
	}

	var sub Subscription
	var response SubscriptionResponse

	row := db.QueryRow("SELECT subsID FROM membership WHERE user_id = ?", user.Id)
	err = row.Scan(&sub.Id)

	if err != nil {
		_, errQuery := db.Exec("INSERT INTO subscription(user_id,type,duration,price) VALUES (?,?,?,?)",
			user.Id,
			jenisloc,
			duration,
			price,
		)

		if errQuery == nil {
			if jenisloc == "Premium" {
				response.Status = 200
				response.Message = "New Premium Subs Added!"
			} else if jenisloc == "Basic" {
				response.Status = 200
				response.Message = "New Basic Subs Added!"
			}
		} else {
			response.Status = 400
			response.Message = "Failed to Register New Subs!"
		}

		row2 := db.QueryRow("SELECT id FROM subscription WHERE user_id = ?", user.Id)
		err2 := row2.Scan(&sub.Id)
		fmt.Println(err2)

		_, errQuery3 := db.Exec("INSERT INTO membership(id,user_id,subsID,CCNumber,validThru,CVC) VALUES (?,?,?,?,?,?)",
			id,
			user.Id,
			sub.Id,
			noKartuKredit,
			masaBerlaku,
			kodeCVC,
		)

		if errQuery3 == nil {
			response.Status = 200
			response.Message = "New Membership Added"
		} else {
			response.Status = 400
			response.Message = "Failed to Register New Membership!"
		}

	} else {
		response.Status = 400
		response.Message = "Already Subscribed!"
	}

	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DeleteSubscription(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	userId := user.Id

	_, errQuery := db.Exec("DELETE FROM membership WHERE user_id = ?",
		userId,
	)

	_, errQuery2 := db.Exec("DELETE FROM subscription WHERE user_id = ?",
		userId,
	)

	var response SubscriptionResponse

	if errQuery == nil && errQuery2 == nil {
		response.Status = 200
		response.Message = "Unsubscribe Success!"
	} else {
		response.Status = 400
		response.Message = "Unsubscribe Failed!"
	}

	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
