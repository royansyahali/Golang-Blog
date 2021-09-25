package request

import "strings"

type TagRequest struct {
	Name   string `json:"name" validate:"required,ascii,max=50"`
	UserId int    `json:"user_id" validate:"required,numeric"`
}

func (t *TagRequest) Prepare() {
	t.Name = strings.ToLower(t.Name)
}
