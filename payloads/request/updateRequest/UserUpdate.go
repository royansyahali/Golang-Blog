package updaterequest

import (
	"time"
)

type UserUpdate struct {
	Name     string `json:"name"  validate:"ascii"`
	Password string `json:"password" validate:"alphanum,min=8"`
	UpdateAt string
}

func (u *UserUpdate) Prepare() {
	u.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}
