package response

type PostResponse struct {
	Id       int      `json:"id"`
	Name     string   `json:"user_name"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	CreateAt string   `json:"create_at"`
	UpdateAt string   `json:"update_at"`
	Category []string `json:"category"`
	Tag      []string `json:"tag"`
}
