package entities

type PostHasTag struct {
	Id     int `json:"id"`
	PostId int `json:"post_id"`
	TagId  int `json:"tag_id"`
}
