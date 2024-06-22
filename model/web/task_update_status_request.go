package web

type TaskUpdateStatusRequest struct {
	Id     int    `validate:"required" json:"id"`
	Status string `validate: "required" json:"status"`
}
