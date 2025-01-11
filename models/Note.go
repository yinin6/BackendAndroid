package models

type Note struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ImageBase64 string `json:"imageBase64"`
	Username    string `json:"username"`
}
