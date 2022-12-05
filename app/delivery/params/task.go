package params

type TaskCreate struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
}
