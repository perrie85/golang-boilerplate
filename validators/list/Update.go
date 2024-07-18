package list

type ListUpdate struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	UserId      int    `validate:"required" json:"user_id"`
}
