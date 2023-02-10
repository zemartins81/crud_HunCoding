package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Age      int8   `json:"age" binding:"required,numeric,min=18,max=100"`
}
