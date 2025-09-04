package vos

type UpdatePostsVO struct {
	Id      string `json:"id" binding:"numeric"`
	Title   string `json:"title" validate:"require,min=1,max=200"`
	Content string `json:"content" validate:"required,min=1,max=10200"`
	State   int    `json:"state"`
}
