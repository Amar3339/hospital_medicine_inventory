package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func RegisterHandler(c *gin.Context){
var req RegisterRequest //RegisterRequest struct data or regustration data are saved in req variable 
err:=c.ShouldBindJSON(&req)
if err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{ 
		"message":"invalid reg data",
	})
	return
}
err=RegisterUser(req)
if err!=nil{
	c.JSON(400,gin.H{"message":"failed to register"})
	return 
}
c.JSON(http.StatusOK,gin.H{
	"message":"sucessfully register datat match",
	"emal":req.Email,
	"username":req.Username,
})
}