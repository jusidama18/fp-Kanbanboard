package params

type TaskCreate struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
}

type TaskPutByID struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type TaskUpdateStatus struct {
	Status bool `json:"status" validate:"required"`
}

type TaskUpdateCategory struct {
	CategoryID int `json:"category_id" validate:"required"`
}
