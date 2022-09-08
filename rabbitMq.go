package main

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"myProject/logics"
	"myProject/repositories"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Request struct {
	Id         string `json:"id"`
	DiscountId string `json:"discountId"`
	Username   string `json:"username"`
	Amount     int    `json:"amount"`
}

func rabbitConsume(channel amqp.Channel) {
	msgs, err := channel.Consume(
		"transactionQueue", // queue
		"",                 // consumer
		false,              // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	failOnError(err, "Failed to register a consumer")

	//var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var test string
			test = string(d.Body)
			var req Request
			json.Unmarshal([]byte(test), &req)

			var wallet repositories.Wallet
			var walletTransaction repositories.WalletTransaction
			wallet.UserName = req.Username

			walletTransaction.Amount = req.Amount
			walletTransaction.Description = "GIFT"
			walletTransaction.Type = 1
			result, _ := logics.AddTransactionLogic(wallet, walletTransaction)
			failOnError(err, "org error")
			fmt.Println("transaction added successfully! =>", result)
			d.Ack(false)
		}
	}()

	//<-forever

}

func rabbitClient() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer channel.Close()

	_, err = channel.QueueDeclare(
		"transactionQueue", // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return channel
}
