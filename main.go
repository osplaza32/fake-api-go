package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"osplaza32/ApiFastToTest/config"
	"osplaza32/ApiFastToTest/controller"
)


func main() {
	port := os.Getenv("PORT")

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	if port == "" {
		port = "8000"
	}
	v1 := router.Group("/api/v1/swagger")
	{
		//v1.POST("/data/:id/fake",controllerfake.PostToThisData)
		v1.GET("/data/:id/fake",controllerfake.GetThisUser)
		v1.GET("/datas/fakes",controllerfake.GetUsers)
		//v1.PUT("/data/:id/fake",controllerfake.EditThisUser )
		//v1.DELETE("/data/:id/fake", )
		}

	authMiddleware,err := Config.MakeJWT("oscar-key")
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", controllerfake.HelloHandler)
		auth.POST("/data/:id/fake",controllerfake.PostToThisData)
	}


	if err := http.ListenAndServe(":"+port,router); err != nil {
		log.Fatal(err)
	}
}
