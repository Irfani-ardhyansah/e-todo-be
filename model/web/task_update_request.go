package web

type TaskUpdateRequest struct {
	Id          int    `validate:"required" json:"id"`
	Name        string `validate:"required" json:"name"`
	Status      string `validate: "required" json:"status"`
	Code        string `validate: "required" json:"code"`
	Description string `validate: "required" json:"description"`
}
