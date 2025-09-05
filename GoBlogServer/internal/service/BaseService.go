package service

type BaseService interface {
}

func CalculateOffsetNum(pageIndex *int, pageSize *int) int {
	if *pageIndex <= 0 {
		*pageIndex = 1
	}
	if *pageSize <= 0 {
		*pageSize = 10
	}
	return (*pageIndex - 1) * (*pageSize)
}
