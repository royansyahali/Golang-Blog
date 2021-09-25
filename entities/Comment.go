package entities

type Comment struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	PostId   int    `json:"post_id"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
