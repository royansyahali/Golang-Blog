package entities

type Post struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
