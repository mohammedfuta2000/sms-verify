package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammedfuta2000/sms-verify/api"
)

func main()  {
	router:= gin.Default()

	// initialize config
	app:=api.Config{Router: router}


	app.Routes()

	router.Run(":8000")

}