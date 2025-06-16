package web

type CommentCreateRequest struct {
	TaskId   int    `validate: "required" json:"task_id"`
	UserId   int    `validate: "required" json:"user_id"`
	ParentId int    `validate: "required" json:"parent_id"`
	Comment  string `validate: "required" json:"comment"`
}
