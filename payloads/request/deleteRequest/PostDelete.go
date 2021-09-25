package deleterequest

type PostDelete struct {
	UserId int `json:"user_id" validate:"required,numeric"`
}
