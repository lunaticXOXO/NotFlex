package controllers

type MovieHistory struct {
	Id       int    `form : "id" json : "id"`
	Title    string `form : "title" json : "title"`
	Director string `form : "director" json : "director"`
}

type MovieHistoryResponse struct {
	Status  int            `form : "status" json : "status"`
	Message string         `form : "message" json : "message"`
	Data    []MovieHistory `form : "data" json : "data"`
}
