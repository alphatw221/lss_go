package rabbitmq

import (
	"fmt"
	"os"
	"strings"

	"lss_go/lib/utils"
	_ "github.com/joho/godotenv/autoload"
	amqp "github.com/rabbitmq/amqp091-go"
	
)


var (
	conn *amqp.Connection
	ch   *amqp.Channel
)

// func init(){

// 	user := os.Getenv("RABBITMQ_USER")
// 	password := os.Getenv("RABBITMQ_PASSWORD")
// 	host := os.Getenv("RABBITMQ_HOST")
// 	port := os.Getenv("RABBITMQ_PORT")

// 	connection_string := fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port)

// 	// 連接RabbitMQ伺服器
// 	var err error
// 	conn, err = amqp.Dial(connection_string)
// 	utils.failOnError(err, "Failed to connect to RabbitMQ")
	

// 	// 建立一個頻道
// 	ch, err = conn.Channel()
// 	utils.failOnError(err, "Failed to open a channel")
	

	

// 	q, err := ch.QueueDeclare(
// 		"stream_tasks", // name
// 		true,   // durable
// 		false,   // delete when unused
// 		false,   // exclusive
// 		false,   // no-wait
// 		nil,     // arguments
// 	  )
// 	utils.failOnError(err, "Failed to declare a queue")
	
	
// 	err = ch.Qos(
// 		1,     // prefetch count
//   		0,     // prefetch size
//   		false, // global
// 	)
// 	utils.failOnError(err, "Failed to set QoS")
	
// }

func GetChannel() *amqp.Channel {
	return ch
}

func Connect(){
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	connection_string := fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port)

	// 連接RabbitMQ伺服器
	var err error
	conn, err = amqp.Dial(connection_string)
	utils.failOnError(err, "Failed to connect to RabbitMQ")
	

	// 建立一個頻道
	ch, err = conn.Channel()
	utils.failOnError(err, "Failed to open a channel")
}


func ListenOn(queueName string, handle func){

	q, err := ch.QueueDeclare(
		queueName, // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )
	utils.failOnError(err, "Failed to declare a queue")
	
	
	err = ch.Qos(
		1,     // prefetch count
  		0,     // prefetch size
  		false, // global
	)
	utils.failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.failOnError(err, "Failed to register a consumer")
	
	var forever chan struct{}
	
	go func() {
		for d := range msgs {

			handle(d.Body)
			
			log.Printf("Done")
			d.Ack(false)
		}
	}()
	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}



func Close(){
	conn.Close()
	ch.Close()
}