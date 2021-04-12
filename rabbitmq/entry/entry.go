package main

import (
	_ "rabbitmq/queue"
)

func main() {
	c := make(chan struct{})
	<-c
}
