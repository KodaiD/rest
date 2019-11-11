package entity

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}