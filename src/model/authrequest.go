package model

type AuthRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
