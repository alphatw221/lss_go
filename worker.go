package main

import (
	// "fmt"
	// "log"
	// "time"
	// "os"

    _ "github.com/joho/godotenv/autoload"

	db "lss_go/lib/db"

	// "lss_go/test"

	// "lss_go/test/test2"

	// _ "lss_go/lib/utils"

)





func main() {

	// godotenv.Load()
    

	// nodesStr := os.Getenv("SCYLLA_NODES")
	// fmt.Println(nodesStr)

	// test2.SayHello()

	// message := test.SayHello()
  	// fmt.Println(message)


	db.Connect()
	defer db.Close()

	// mq.Connect()
	// defer mq.Close()



}