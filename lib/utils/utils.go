package utils

import (
	"fmt"

)


func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println("%s: %s", msg, err)
		// log.Panicf("%s: %s", msg, err)
	}
}

func test(t string){
	fmt.Println("%", t)

}