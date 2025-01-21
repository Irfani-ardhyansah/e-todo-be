package web

type TimerResponse struct {
	Id        int    `json:"id"`
	TaskId    int    `json:"task_id"`
	Time      string `json:"time"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
