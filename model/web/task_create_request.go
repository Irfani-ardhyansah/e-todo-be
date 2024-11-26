package web

type TaskCreateRequest struct {
	Name   string `validate: "required" json:"name"`
	Status string `validate: "required" json:"status"`
	Code   string `validate: "required" json:"code"`
}
