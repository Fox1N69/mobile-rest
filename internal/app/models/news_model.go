package models

type NewsData struct {
	ID       uint          `json:"id" gorm:"primaryKey"`
	Title    string        `json:"title"`
	Content  string        `json:"content"`
	Link     string        `json:"link"`
	ImageUrl string        `json:"image_url"`
	FullNews  FullNewsData   `json:"full_news"` 
}

type FullNewsData struct {
	ID           uint   `gorm:"primaryKey"`
	NewsDataID   uint   `json:"newsdata_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	NewsImageUrl string `json:"newsimage_url"`
	Link         string `json:"link"`
}


func (news *NewsData) SetFullNews(fullNews FullNewsData) {
	news.FullNews = fullNews

}