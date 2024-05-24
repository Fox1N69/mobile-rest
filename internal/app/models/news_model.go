package models

type NewsData struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Link     string `json:"link"`
	ImageUrl string `json:"image_url"`
	FullNews *FullNewsData `gorm:"foreignKey:NewsDataID"`
}

type FullNewsData struct {
	ID uint `gorm:"primaryKey"`
	NewsDataID uint   `json:"newsdata_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Link string `json:"link"`
}

func (f *FullNewsData) SetTitleFromNewsData(news *NewsData) {
	f.Title = news.Title
}
