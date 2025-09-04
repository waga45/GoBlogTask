package vos

type GetPostListVO struct {
	Title     string `json:"title"`
	State     int    `json:"state"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}
