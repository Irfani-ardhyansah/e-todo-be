package web

type CommentResponse struct {
	Id        int                 `json:"id"`
	TaskId    int                 `json:"task_id"`
	UserId    int                 `json:"user_id"`
	ParentId  *int                `json:"parent_id"`
	Comment   string              `json:"message"`
	CreatedAt string              `json:"time"`
	User      CommentUserResponse `json:"user"`
	Childs    []*CommentResponse  `json:"response"`
}

type CommentUserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
