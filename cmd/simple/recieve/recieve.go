package main

import rabbitmq2 "rabbitmq-demo"

func main() {
	rabbitmqc := rabbitmq2.NewRabbitMQSimple("mqtest01")
	rabbitmqc.ConsumeSimple()
}
