package web

type TimerCreateRequest struct {
	TaskId    int    `validate: "required" json:"task_id"`
	Time      string `validate: "required" json:"time"`
	Title     string `json:"title"`
	Status    string `validate: "required" json:"status"`
	CreatedAt string `validate: "required" json:"created_at"`
}
