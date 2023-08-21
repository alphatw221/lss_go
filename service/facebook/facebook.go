package facebook

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"

	// "time"

	types "lss_go/lib/types"
	http "lss_go/service/http"
)

// type PictureData struct {
// 	// Height       int    `json:"height"`
// 	// Width        int    `json:"width"`
// 	// IsSilhouette bool   `json:"is_silhouette"`
// 	Url string `json:"url"`
// }
// type Picture struct {
// 	Data PictureData `json:"data"`
// }
// type From struct {
// 	Picture Picture `json:"picture"`
// 	Name    string  `json:"name"`
// 	ID      string  `json:"id"`
// }

// type CommentData struct {
// 	CreatedTime int    `json:"created_time"`
// 	From        From   `json:"from"`
// 	Message     string `json:"message"`
// 	ID          string `json:"id"`
// }

// type ErrorData struct {
// 	Message string `json:"message"`
// 	Type    string `json:"type"`
// 	Code    int    `json:"code"`
// }

// type CommentResData struct {
// 	Error ErrorData     `json:"error"`
// 	Data  []CommentData `json:"data"`
// }

func GetPostComments(page_token string, post_id string, since int) types.CommentResData {

	params := map[string]string{
		"since": strconv.Itoa(since), "order": "chronological", "limit": "100",
		"date_format": "U", "live_filter": "no_filter",
		"fields": "created_time,from{picture,name,id},message,id",
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", page_token),
	}

	path := fmt.Sprintf("%s/comments", post_id)

	domain := "graph.facebook.com"
	body := http.Get(domain, path, params, headers)

	var resData types.CommentResData
	err := json.Unmarshal(body, &resData)

	if err != nil {
		fmt.Println("JSON解碼錯誤:", err)
	}

	return resData

}

// type LiveCommentResData struct {
// 	Data CommentData `json:"data"`
// }

func StreamLiveComments(page_token string, live_video_id string, ch chan<- types.LiveCommentResData) {

	domain := "streaming-graph.facebook.com"

	params := map[string]string{
		"access_token": page_token,
		"comment_rate": "one_per_two_seconds",
		"fields":       "created_time,from{picture,name,id},message,id",
	}

	path := fmt.Sprintf("%s/live_comments", live_video_id)

	res := http.GetStreamEvent(domain, path, params, nil)

	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)

	for {
		line, err := reader.ReadString('\n')

		if err == nil {
			line := []byte(line)
			if string(line[0:4]) == "data" {
				line = line[6:]
				fmt.Println(line)
				var liveResData types.LiveCommentResData
				json.Unmarshal(line, &liveResData)

				ch <- liveResData
			}

			// fmt.Println("收到事件:", line)
			// fmt.Println("收到事件:", line)

		}

	}

}
