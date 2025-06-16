package web

type CommentResponse struct {
	Id        int    `json:"id"`
	TaskId    int    `json:"task_id"`
	UserId    int    `json:"user_id"`
	ParentId  int    `json:"parent_id"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}
