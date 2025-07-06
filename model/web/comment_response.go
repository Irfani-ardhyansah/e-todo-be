package web

type CommentResponse struct {
	Id        int                `json:"id"`
	TaskId    int                `json:"task_id"`
	UserId    int                `json:"user_id"`
	ParentId  *int               `json:"parent_id"`
	Comment   string             `json:"message"`
	CreatedAt string             `json:"time"`
	Childs    []*CommentResponse `json:"response"`
}
