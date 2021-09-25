package request

import "time"

type CommentRequest struct {
	UserId   int    `json:"user_id" validate:"required,numeric"`
	PostId   int    `json:"post_id" validate:"required,numeric"`
	Content  string `json:"content" validate:"required"`
	CreateAt string
	UpdateAt string
}

func (c *CommentRequest) Prepare() {
	c.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	c.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}
