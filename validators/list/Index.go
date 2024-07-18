package list

type ListIndex struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	UserId      int    `validate:"required" json:"user_id"`
	CreatedAt   string `validate:"required,datetime" json:"created_at"`
	OrderBy     *orderBy
}

type orderBy struct {
	By        string
	Direction string
}
