package request

import "strings"

type CategoryRequest struct {
	Name   string `json:"name" validate:"required,ascii,max=50"`
	UserId int    `json:"user_id" validate:"required,numeric"`
}

func (c *CategoryRequest) Prepare() {
	c.Name = strings.ToLower(c.Name)
}
