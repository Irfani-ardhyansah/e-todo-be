package web

type TimerUpdateRequest struct {
	Id     int    `validate: "required" json:"id"`
	Time   string `validate: "required" json:"time"`
	Title  string `json: "title"`
	Status string `validate: "required" json:"status"`
}
