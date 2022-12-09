package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

// Receiver 消息接收實體
type Receiver struct {
	name      string
	queueName string
	conn      *amqp.Connection
	channel   *amqp.Channel
	handlers  map[string]func([]byte) // consumer message handler
	notify    chan *amqp.Error        // disconnect notify
	quit      chan error
}

var wg sync.WaitGroup

var historyHooks = []func(gameID, num string){}

// SetHistoryHook 設定賽果接收通知 hook
func SetHistoryHook(fn func(gameID, num string)) {
	historyHooks = append(historyHooks, fn)
}

func triggerHistoryHook(gameID, num string) {
	for _, fn := range historyHooks {
		fn(gameID, num)
	}
}

func main() {
	wg.Add(2)
	i := &Receiver{
		name:      "hawkeye",
		queueName: "risk_close",
		handlers: map[string]func([]byte){
			"close": PreCloseReceiver,
		},
	}

	i.Connect()

	go i.ReConnect()
	// go i.GoClose()

	wg.Wait()
}

func (i *Receiver) Connect() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return
	}

	i.conn = conn
	i.channel = channel
	// if err := i.channel.ExchangeDeclare(
	// 	i.name,
	// 	amqp.ExchangeFanout,
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// ); err != nil {
	// ilog.Log("receiver", fn).AddParams(struct {
	// 	Name           string
	// 	ExchangeFanout string
	// 	Err            string
	// }{i.name, amqp.ExchangeFanout, err.Error()}, errs.ErrExternal.Err()).Error("i.channel.ExchangeDeclare()")
	// return
	// }

	if _, err = i.channel.QueueDeclare(i.queueName, false, false, false, false, nil); err != nil {
		fmt.Println(err)
		// ilog.Log("receiver", fn).AddParams(struct {
		// 	QueueName string
		// 	Err       string
		// }{i.queueName, err.Error()}, errs.ErrExternal.Err()).Error("i.channel.QueueDeclare()")
		return
	}

	// if err := i.channel.QueueBind(i.queueName, "", i.name, false, nil); err != nil {
	// ilog.Log("receiver", fn).AddParams(struct {
	// 	QueueName string
	// 	Err       string
	// }{i.queueName, err.Error()}, errs.ErrExternal.Err()).Error("i.channel.QueueBind()")
	// return
	// 	return
	// }

	// ----- Consumer -----
	msgs, err := i.channel.Consume(i.queueName, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		// ilog.Log("receiver", fn).AddParams(i.queueName, errs.ErrExternal.Err()).Error("i.channel.Consume()")

		return
	}
	go i.Watch(msgs)

	i.notify = i.conn.NotifyClose(make(chan *amqp.Error))
}

// Watch 監聽訊息
func (i *Receiver) Watch(delivery <-chan amqp.Delivery) {

	for d := range delivery {
		// PreCloseReceiver(d.Body)
		if fn, ok := i.handlers["close"]; ok {
			fn(d.Body)
		}
	}
}

// ReConnect 斷線重連
func (i *Receiver) ReConnect() {
	// fn := "Receiver.ReConnect"

	for {
		select {
		case err := <-i.notify:
			fmt.Println("有進來嘛A")
			if err != nil {
				// ilog.Log("receiver", fn).AddParams(err, errs.ErrExternal.Err()).Error("rabbitmq consumer")
				panic(err)
			}
		case <-i.quit:
			fmt.Println("有進來嘛B")
			return
		}

		if !i.conn.IsClosed() {
			fmt.Println("有進來嘛C")
			if err := i.conn.Close(); err != nil {
				panic(err)
			}
		}

	quit:
		for {
			select {
			case <-i.quit:
				fmt.Println("有進來嘛D")
				return

			default:
				fmt.Println("有進來嘛E")
				i.Connect()
				// if err := i.Connect(); err != nil {
				// 	panic(err)
				// 	// sleep 5s reconnect
				// 	time.Sleep(time.Second * 5)
				// 	continue
				// }

				break quit
			}
		}
	}
}

func (i *Receiver) GoClose() {
	time.Sleep(10 * time.Second)
	i.conn.Close()
	fmt.Println("已close")
}

func PreCloseReceiver(b []byte) {
	fmt.Println("PreCloseReceiver => in")
	data := struct {
		GameId string `json:"game_id"`
		Num    string `json:"num"`
	}{}
	if err := json.Unmarshal(b, &data); err != nil {
		return
	}
	fmt.Println("PreCloseReceiver => ", data.GameId)
	time.Sleep(61 * time.Second)
	fmt.Println("應該要timeout 出去喔")
}
