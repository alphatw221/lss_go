package mq

import (
	rabbitmq "lss_go/lib/mq/rabbitmq"
)

func Connect() {
	rabbitmq.Connect()
}

func Close() {
	rabbitmq.Close()
}
