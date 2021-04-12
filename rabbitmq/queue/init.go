package queue

import (
	"encoding/json"
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
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//failOnError(err, "Failed to connect to RabbitMQ")
	//
	//ch, err := conn.Channel()
	//failOnError(err, "Failed to open a channel")
	//
	//AddConsumer(ch, "queue_1", Handler1)
	//AddConsumer(ch, "queue_2", Handler2)
	//AddConsumer(ch, "queue_3", Handler3)

	InitAndConsumeQueue("queue_1", Handler1)
	InitAndConsumeQueue("queue_2", Handler2)
	InitAndConsumeQueue("queue_3", Handler3)
}

//func AddConsumer(ch *amqp.Channel, queue string, f func(msg *amqp.Delivery)) {
//	msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
//	if err != nil {
//		failOnError(err, "Failed to consume queue")
//		return
//	}
//	go func() {
//		for {
//			select {
//			case msg := <-msgs:
//				f(&msg)
//			}
//		}
//	}()
//
//	log.Printf("Init %v successful", queue)
//}

func Handler1(msg *amqp.Delivery) {

	info := vo.QueueMsg{}
	if err := json.Unmarshal(msg.Body, &info); err != nil {
		failOnError(err, "Unmarshal msg failed.")
		return
	}

	log.Printf("Handler 1 msg : %v", info)
}

func Handler2(msg *amqp.Delivery) {

	info := vo.QueueMsg{}
	if err := json.Unmarshal(msg.Body, &info); err != nil {
		failOnError(err, "Unmarshal msg failed.")
		return
	}

	log.Printf("Handler 2 msg : %v", info)
}

func Handler3(msg *amqp.Delivery) {
	info := vo.QueueMsg{}
	if err := json.Unmarshal(msg.Body, &info); err != nil {
		failOnError(err, "Unmarshal msg failed.")
		return
	}

	log.Printf("Handler 3 msg : %v", info)
}
