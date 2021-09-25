package entities

type PostHasCategory struct {
	Id         int `json:"id"`
	PostId     int `json:"post_id"`
	CategoryId int `json:"category_id"`
}
