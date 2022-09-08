package main

import (
	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"myProject/db"
	"net/http"
)

var channel *amqp.Channel

func main() {

	db.GetSqlConnection()
	channel = rabbitClient()
	rabbitConsume(*channel)
	log.SetFormatter(&log.JSONFormatter{})
	r := mux.NewRouter()
	r.HandleFunc("/wallets", add).Methods("POST")
	r.HandleFunc("/wallets/{username}/transactions", addTransaction).Methods("POST")
	r.HandleFunc("/wallets/{username}", get).Methods("GET")
	http.ListenAndServe("localhost:8080", r)

}
