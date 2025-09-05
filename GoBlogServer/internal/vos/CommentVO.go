package vos

type CommentVO struct {
	Id     int    `json:"id"`
	PostId string `json:"postId" binding:"required,numeric,min=1"`
}

// 添加评论bean
type PushCommentVO struct {
	CommentVO
	Content string `json:"content"`
}

// 获取列表vo
type GetCommentListVO struct {
	CommentVO
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}
