package web

type CommentUpdateRequest struct {
	Id      int    `validate:"required" json:"id"`
	Comment string `validate: "required" json:"comment"`
}
