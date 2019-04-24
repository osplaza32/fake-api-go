package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"osplaza32/ApiFastToTest/controller"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())


	v1 := router.Group("/api/v1/swagger")
	{
		v1.POST("/data/:id/fake",controllerfake.PostToThisData)
		v1.GET("/data/:id/fake",controllerfake.GetThisUser)
		v1.GET("/datas/fakes",controllerfake.GetUsers)
		//v1.PUT("/data/:id/fake", )
		//v1.DELETE("/data/:id/fake", )


	}


	router.Run(":8000")
	
}
