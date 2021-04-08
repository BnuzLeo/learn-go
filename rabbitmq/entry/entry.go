package main

import (
	"log"
	_ "rabbitmq/queue"
)

func main() {
	c := make(chan struct{})
	log.Printf("Start to study rabbit mq...")
	<-c
}
