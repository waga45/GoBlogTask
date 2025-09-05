package dtos

type BasePageDto[T any] struct {
	TotalCount int64 `json:"totalCount"`
	List       *[]T  `json:"list"`
}
