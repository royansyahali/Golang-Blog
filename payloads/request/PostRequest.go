package request

import (
	"time"
)

type PostRequest struct {
	UserId     int    `json:"user_id" validate:"required,numeric"`
	Title      string `json:"title" validate:"required,max=50"`
	Content    string `json:"content" validate:"required"`
	CategoryId string `json:"category_id" validate:"required,ascii"`
	TagId      string `json:"tag_id" validate:"required,ascii"`
	CreateAt   string
	UpdateAt   string
}

func (p *PostRequest) Prepare() {
	p.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	p.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

}
