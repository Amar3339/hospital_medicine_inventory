package login

import (

	"net/http"

	"github.com/gin-gonic/gin"
)
func LoginUser(c *gin.Context){
	var log Login
	err:=c.ShouldBindJSON(&log)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"failed to login"})
		
	
return
	}
err=LoginDetails(log)
if err!=nil{
	c.JSON(400, gin.H{
		"message":"failed to login "	})
		

	return
}
c.JSON(http.StatusOK,gin.H{
	"message":"login suceefully",
	"email":log.Email,
	"username":log.Username,
})
}
	
