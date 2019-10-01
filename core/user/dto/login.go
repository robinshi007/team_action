package dto

// IdentityKey -
var IdentityKey = "userid"

// IdentityName -
var IdentityName = "username"

// Login -
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
