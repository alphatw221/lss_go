package types

type PictureData struct {
	// Height       int    `json:"height"`
	// Width        int    `json:"width"`
	// IsSilhouette bool   `json:"is_silhouette"`
	Url string `json:"url"`
}
type Picture struct {
	Data PictureData `json:"data"`
}
type From struct {
	Picture Picture `json:"picture"`
	Name    string  `json:"name"`
	ID      string  `json:"id"`
}

type CommentData struct {
	CreatedTime int    `json:"created_time"`
	From        From   `json:"from"`
	Message     string `json:"message"`
	ID          string `json:"id"`
}

type ErrorData struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    int    `json:"code"`
}

type CommentResData struct {
	Error ErrorData     `json:"error"`
	Data  []CommentData `json:"data"`
}

type LiveCommentResData struct {
	Data CommentData `json:"data"`
}
