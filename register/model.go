package register

import (
	
)
type registerRequest struct{
	Business_name string `json:"business_name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Address string `json:"address" binding:"required"`
	City string `json:"city"`
	Email string `json:"email" binding:"required"`
	PhoneNumber string  `json:"phone_number"`//(because pphone number have leading 0 and it contains + and this is not part of integer)// 
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	


}