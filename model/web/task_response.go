package web

type TaskResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
}
