package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// test
	conn, err := amqp.Dial("amqp://root:abcd1234!@localhost:5672/")
	if err != nil {
		fmt.Printf("%s: %s\n", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"history", // queue
		"",        // consumer
		true,      // Auto Ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	type List struct {
		GameId string `json:"game_id"`
		Num    string `json:"num"`
	}
	go func() {
		for d := range msgs {
			list := &List{}
			json.Unmarshal([]byte(d.Body), &list)
			fmt.Println("時間：" + time.Now().Format("2006-01-02 15:04:05") + "\t| 遊戲名稱：" + list.GameId + "\t| 期數：" + list.Num)
		}
	}()
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
