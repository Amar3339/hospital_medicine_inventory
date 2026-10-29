package main

import (
	"medical_inventory/login"
	"medical_inventory/register"

	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {
	register.RegisterRouter(router)
	login.Logged(router)

}
