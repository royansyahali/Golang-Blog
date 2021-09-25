package entities

type User struct {
	Id       int    `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
