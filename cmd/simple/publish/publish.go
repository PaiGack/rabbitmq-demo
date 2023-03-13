package main

import (
	"fmt"
	rabbitmq2 "rabbitmq-demo"
)

func main() {
	rabbitmqc := rabbitmq2.NewRabbitMQSimple("mqtest01")
	for i := 0; i < 1024; i++ {
		rabbitmqc.PublishSimple(fmt.Sprintf("Hello mqtest %d!", i))
		fmt.Println("send success ", i)
	}
}
