package login

type Login struct {
	BusinessName string `json:"business_name" binding:"required"`
	Username      string `json:"username" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}







