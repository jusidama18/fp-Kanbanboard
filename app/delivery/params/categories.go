package params

type CategoryCreate struct {
	Type string `json:"type" validate:"required"`
}

type CategoryUpdate struct {
	Type string `json:"type" validate:"required"`
}
