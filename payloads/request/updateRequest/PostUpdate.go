package updaterequest

import (
	"time"
)

type PostUpdate struct {
	UserId     int    `json:"user_id" validate:"required,numeric"`
	Title      string `json:"title" validate:"max=50"`
	Content    string `json:"content"`
	CategoryId string `json:"category_id" validate:"ascii"`
	TagId      string `json:"tag_id" validate:"ascii"`
	UpdateAt   string
}

func (u *PostUpdate) Prepare() {
	u.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}
