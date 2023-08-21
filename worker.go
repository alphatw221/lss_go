package main

import (
	"fmt"
	// "log"
	// "time"
	// "os"

	_ "github.com/joho/godotenv/autoload"

	// db "lss_go/lib/db"
	// mq "lss_go/lib/mq"

	types "lss_go/lib/types"
	facebook "lss_go/service/facebook"
)

func main() {

	// db.Connect()
	// defer db.Close()

	// mq.Connect()
	// defer mq.Close()

	page_token := "EAANwBngXqOABANDaDPEVmVIhZBchXZAMH9loiJE3L5sjCnlBPucrZArXSukVvIgXcZCZAZATE7t4fZChnW9BGmUZA5zZAtnMtchsVr5msMa4wZA6RFBtKyZBCOSyMrzwUC51TQQz0LQ6MRCNx0hPV7rj5ZABEN5i7ZB4wDZBgrKpkEKZB0ZCyp2ZBZCemkRaTVFWffBGVvnLBdwA63I1vGOAZDZD"
	// facebook.GetPostComments(
	// 	page_token,         //page_token
	// 	"3380256538951955", //post_id
	// 	1,                  //since
	// )

	ch := make(chan types.LiveCommentResData)

	go facebook.StreamLiveComments(
		page_token,
		"3380256538951955",
		ch,
	)

	for {
		comment := <-ch
		fmt.Println("testtest")

		fmt.Println(comment)
	}
}
