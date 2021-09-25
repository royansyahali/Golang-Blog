package request

import (
	"strings"
	"time"
)

type UserRequest struct {
	Name     string `json:"name"  validate:"required,ascii"`
	Username string `json:"username" validate:"required,alpha"`
	Password string `json:"password" validate:"required,alphanum,min=8"`
	Role     string
	CreateAt string
	UpdateAt string
}

func (u *UserRequest) Prepare() {
	u.Role = "guest"
	u.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	u.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	u.Username = strings.ToLower(u.Username)
}
