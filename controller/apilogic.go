package controllerfake
import (
	"github.com/gin-gonic/gin"
	"github.com/wawandco/fako"
	"net/http"
	"strconv"
)
type User struct {
	Id int
	Name     string `fako:"full_name"`
	Username string `fako:"user_name"`
	Email    string `fako:"email_address"`//Notice the fako:"email_address" tag
	Phone    string `fako:"phone"`
	Password string `fako:"simple_password"`
	Address  string `fako:"street_address"`
	Extras DataUser
 }
type DataUser struct {
	Id int
	Edad     string
	Compania string
	ColorOjos    string
	ColorPelo    string

}
type Users struct {
	app string
	users []User
}
func (u Users) makeuser() Users {
	for i := 0; i < 10; i++ {
		u.users = append(u.users, makeUser())
		u.users[i].Id = i
		}
	return u
	}
func (u Users) checkparameterFill(i int)bool {
	if  len(u.users)-1 >= i {
		return true
		}
	return false
}
func makeUser() User {
	var user User
	fako.Fill(&user)
	return user

}
func GetUsers(c *gin.Context){
	var users Users
	users.app = "FaKE aPi"
	users = users.makeuser()
	c.JSON(http.StatusOK,gin.H{"algo":users.app,"data":users.users})
	}
func PostToThisData(c *gin.Context){


}
func GetThisUser(c *gin.Context) {
	var users Users
	users.app = "FaKE aPi"
	users = users.makeuser()
	userid := c.Param("id")
	i1, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Message":"El elemento enviado en el parametro no es un numero"})
		return
		}
	if users.checkparameterFill(i1-1) == false{
		c.JSON(http.StatusNotFound,gin.H{"Message":"El elemento no encontrado"})
		return
		}
	c.JSON(http.StatusOK,gin.H{"algo":users.app,"data":users.users[i1-1]})
}