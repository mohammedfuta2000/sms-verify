package api

import (
	"github.com/gin-gonic/gin"
)

type Config struct{
	Router *gin.Engine
}

func (c *Config) Routes(){
	c.Router.GET("/v1",c.check())
	c.Router.POST("/otp", c.sendSMS())
	c.Router.POST("/verifyotp",c.verifySMS() )
}