package web

type TaskResponse struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
