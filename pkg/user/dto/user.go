package dto

// User -
type NewUser struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=4,lte=20"`
}

// EditPasswordUser -
type EditPasswordUser struct {
	Password string `json:"password" binding:"required,gte=4,lte=20"`
}
