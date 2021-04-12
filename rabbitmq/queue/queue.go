package queue

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

const DEFAULT_RABBIT_MQ = "amqp://guest:guest@localhost:5672/"

type QueueConsumer struct {
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string // 队列名称
	Exchange  string // 交换机
	Key       string // Key
	URL       string // 连接信息
	Retry     bool   // true重新尝试再次连接rabbitmq
}

func InitAndConsumeQueue(queue string, f func(msg *amqp.Delivery)) {
	go func() {
		for {
			log.Printf("Start to init consumer to consume %v", queue)
			consumer, err := NewSimpleQueueConsumer(queue)
			if err != nil {
				if consumer.Retry {
					time.Sleep(time.Second * 5)
					continue
				} else {
					log.Printf("Consumer init failed, queue:%v", queue)
					break
				}
			}
			consumer.Consume(f)
			time.Sleep(time.Second * 5)
		}
	}()
}

func NewQueueConsumer(queue, exchange, key, url string) (consumer *QueueConsumer, err error) {
	consumer = &QueueConsumer{
		QueueName: queue,
		Exchange:  exchange,
		Key:       key,
		URL:       url,
		Retry:     true,
	}

	if consumer.QueueName == "" {
		log.Printf("Queue can't be empty")
		err = errors.New("queue can't be empty")
		consumer.Retry = false
		return
	}

	if consumer.URL == "" {
		log.Printf("Url can't be empty")
		err = errors.New("Url can't be empty")
		consumer.Retry = false
		return
	}

	// 创建rabbitmq连接
	consumer.Conn, err = amqp.Dial(consumer.URL)
	if err != nil {
		log.Printf("Rabbit mq connect failed, url: %v", consumer.URL)
		return
	}

	consumer.Channel, err = consumer.Conn.Channel()
	if err != nil {
		log.Printf("Rabbit mq create channel failed, queue: %v", consumer.QueueName)
		return
	}

	return
}

func NewSimpleQueueConsumer(queue string) (consumer *QueueConsumer, err error) {
	return NewQueueConsumer(queue, "", "", DEFAULT_RABBIT_MQ)
	//return NewQueueConsumer(queue, "", "", "")
}

func (c *QueueConsumer) Destroy() {
	_ = c.Conn.Close()
	_ = c.Channel.Close()
	log.Printf("Queue %v 's consumer was be destroied.", c.QueueName)
}

func (c *QueueConsumer) Consume(f func(msg *amqp.Delivery)) (err error) {
	msgs, err := c.Channel.Consume(c.QueueName, "", true, false, false, false, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("Consumer %v consume %v failed.", c.Key, c.QueueName))
		return
	}

	forever := make(chan bool)
	go func() {
		for {
			select {
			case msg, ok := <-msgs:
				log.Printf("msg:%v, ok:%v", string(msg.Body), ok)
				if !ok {
					forever <- false
					log.Printf("Consume %v failed.", c.QueueName)
					break
				}
				f(&msg)
			}
		}
	}()
	log.Printf("Consume %v successful.", c.QueueName)
	<-forever
	c.Destroy()
	return
}
