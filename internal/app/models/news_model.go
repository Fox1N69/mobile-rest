package models 

type NewsData struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Link string `json:"link"`
}

type NewsDataFull struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}
