package dto

type UserRequest struct {
	Fullname string `json:"fullname" binding:"required,min=6,max=250"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
