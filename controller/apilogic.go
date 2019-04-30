package controllerfake

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"osplaza32/ApiFastToTest/Models"
)



func GetUsers(c *gin.Context){
	var users Models.Users
	users.App = "FaKE aPi"
	users = users.Makeuser()
	c.JSON(http.StatusOK,gin.H{"algo":users.App,"data":users.Users})
	}
func PostToThisData(c *gin.Context){
	var users Models.Users
	users = users.Makeuser()

	var data Models.DataUser
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Message":"El elemento enviado presenta caracteristicas inusuales"+err.Error()})
		return
	}
	userid := c.Param("id")
	i1, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Message":"El elemento enviado en el parametro no es un numero"})
		return
	}
	if users.CheckparameterFill(i1-1) == false{
		c.JSON(http.StatusNotFound,gin.H{"Message":"El elemento no encontrado"})
		return
	}
	users.Users[i1-1].Extras = data
	c.JSON(http.StatusCreated,gin.H{"Message":"El elemento se creo bien","Data":users.Users[i1-1]})





}
func EditThisUser(c *gin.Context){
	var userEdir  Models.UserEdit
	var users Models.Users
	users = users.Makeuser()

	 userid := c.Param("id")
	i1, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Message":"El elemento enviado en el parametro no es un numero"})
		return
	}
	if users.CheckparameterFill(i1-1) == false{
		c.JSON(http.StatusNotFound,gin.H{"Message":"El elemento no encontrado"})
		return
	}
	err = c.BindJSON(&userEdir)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Message":"El elemento enviado presenta caracteristicas inusuales"+err.Error()})
		return
	}
	returnelemen := SeeDiff(users.Users[i1-1],userEdir)
	//fmt.Println(returnelemen)
	c.JSON(http.StatusOK,gin.H{"data":returnelemen})
	}

func SeeDiff(user Models.UserFake, edit Models.UserEdit) Models.UserFake {
	v := reflect.ValueOf(user)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		fmt.Println(reflect.TypeOf(values[i]))
		}

	return user
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
func GetThisUser(c *gin.Context) {
	var users Models.Users
	users.App = "FaKE aPi"
	users = users.Makeuser()
	userid := c.Param("id")
	i1, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Message":"El elemento enviado en el parametro no es un numero"})
		return
		}
	if users.CheckparameterFill(i1-1) == false{
		c.JSON(http.StatusNotFound,gin.H{"Message":"El elemento no encontrado"})
		return
		}
	c.JSON(http.StatusOK,gin.H{"algo":users.App,"data":users.Users[i1-1]})
}
func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(Models.IdentityKey)


	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*Models.User).UserName,
		"text":     "Hello World.",
	})
}