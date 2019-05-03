package Config

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"osplaza32/ApiFastToTest/Models"
	"osplaza32/ApiFastToTest/Models/modelsdb"
	"time"
)

func MakeJWT(TokenHeadName string)(*jwt.GinJWTMiddleware, error)  {

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: Models.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*modelsdb.User); ok {
				return jwt.MapClaims{
					Models.IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &modelsdb.User{
				ID: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			email := loginVals.Username

			user := modelsdb.User{}
			db,err := Conneccion()
			if err != nil {
				panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
			}
			db.Where(&modelsdb.User{Email:email}).First(&user)
			err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(loginVals.Password))
			if err == nil {
				return &user, nil


			}else{
				return nil, jwt.ErrFailedAuthentication

			}


			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)
			id := claims["id"].(string)
			user := modelsdb.User{}
			db,err := Conneccion()
			if err != nil {
				panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
			}
			db.Where(&modelsdb.User{ID:id}).First(&user)
			fmt.Println(user)
			if user.Email == ""{
				return false

			}

			return true

		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: TokenHeadName,

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})



}