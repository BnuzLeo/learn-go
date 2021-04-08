package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/vo"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func init() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	AddConsumer(ch, "queue_1", Handler1)
	AddConsumer(ch, "queue_2", Handler2)
	AddConsumer(ch, "queue_3", Handler3)
}

func AddConsumer(ch *amqp.Channel, queue string, f func(msg *amqp.Delivery)) {
	msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		failOnError(err, "Failed to consume queue")
		return
	}
	go func() {
		for {
			select {
			case msg := <-msgs:
				f(&msg)
			}
		}
	}()

	log.Printf("Init %v successful", queue)
}

func Handler1(msg *amqp.Delivery) {
	defer Recover(context.Background(), func(e interface{}) string {
		err := e.(error)
		return fmt.Sprintf("Handler1 error: %v", err)
	})

	info := vo.QueueMsg{}
	if err := json.Unmarshal(msg.Body, &info); err != nil {
		failOnError(err, "Unmarshal msg failed.")
		return
	}

	log.Printf("Handler 1 msg : %v", info)
}

func Handler2(msg *amqp.Delivery) {
	defer Recover(context.Background(), func(e interface{}) string {
		err := e.(error)
		return fmt.Sprintf("Handler2 error: %v", err)
	})

	info := vo.QueueMsg{}
	if err := json.Unmarshal(msg.Body, &info); err != nil {
		failOnError(err, "Unmarshal msg failed.")
		return
	}

	log.Printf("Handler 2 msg : %v", info)
}

func Handler3(msg *amqp.Delivery) {
	defer Recover(context.Background(), func(e interface{}) string {
		err := e.(error)
		return fmt.Sprintf("Handler3 error: %v", err)
	})

	info := vo.QueueMsg{}
	if err := json.Unmarshal(msg.Body, &info); err != nil {
		failOnError(err, "Unmarshal msg failed.")
		return
	}

	log.Printf("Handler 3 msg : %v", info)
}

func Recover(ctx context.Context, arg0 interface{}, args ...interface{}) {
	if err := recover(); err != nil {
		switch first := arg0.(type) {
		case func(interface{}) string:
			// the recovered err will pass to this func
			//Critical(arg0, append([]interface{}{err}, args)...)
			failOnError(nil, fmt.Sprintf("%v", first))
		case string:
			//Critical(a+"\n%v", append(args, err)...)
			failOnError(nil, "string")
		default:
			//Critical(arg0, append(args, err)...)
			failOnError(nil, "default")
		}
	}
}
