package dto

type User struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required,gte=4,lte=20"`
}
