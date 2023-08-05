package mq

import (
	rabbitmq "lss_go/lib/rabbitmq"
)


func Connect(){
	rabbitmq.Connect()
}

func Close(){
	rabbitmq.Close()
}