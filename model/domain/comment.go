package domain

type Comment struct {
	Id        int
	TaskId    int
	UserId    int
	ParentId  *int
	Comment   string
	CreatedAt string
	UpdatedAt string
	UserName  string
}
