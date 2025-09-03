package dtos

type LoginDto struct {
	Token    string `json:"token"`
	UserName string `json:"username"`
}
