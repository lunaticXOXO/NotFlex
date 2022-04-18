package controllers

type Subscription struct {
	Id       int     `form : "id" json : "id"`
	Types    string  `form : "types" json : "types"`
	Duration string  `form : "duration" json : "duration"`
	Price    float32 `form : "price" json : "price"`
}

type SubscriptionResponse struct {
	Status  int            `form : "status" json : "status"`
	Message string         `form : "message" json : "message"`
	Data    []Subscription `form : "data" json : "data"`
}
