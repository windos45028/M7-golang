package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

// func htmlWelcome(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	w.Header().Set("Content-Type", "text/html")
// 	html := `<!doctype  html>
//     <META  http-equiv="Content-Type"  content="text/html"  charset="utf-8">
//     <html  lang="zhCN">
//         <head>
//             <title>Golang</title>
//             <meta  name="viewport"  content="width=device-width,  initial-scale=1.0,  maximum-scale=1.0,  user-scalable=0;"  />
//         </head>
//         <body>
//             <div id="app">
// 				Welcome!
// 				<a href="/Login">Login</a>
// 				<br>
// 				<a href="/Panic">list</a>
// 				<br>
// 				<a href="/post">post</a>
// 			</div>
//         </body>
//     </html>`
// 	fmt.Fprintf(w, html)
// }

// func Panic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("哈 天真a吉")
// 		}
// 	}()

// 	panic("test")
// }

func main() {
	// mux := httprouter.New()
	// mux.GET("/Login", htmlWelcome)
	// mux.GET("/Panic", Panic)

	// // Set the parameters for a HTTP server
	// server := http.Server{
	// 	Addr:    "0.0.0.0:8000",
	// 	Handler: mux,
	// }

	// fmt.Println("Http Server :8000")

	// server.ListenAndServe()

	// panic("test")

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
