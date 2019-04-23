package controllerfake

import (
	"github.com/gin-gonic/gin"
	"github.com/wawandco/fako"
	"net/http"
)

type User struct {
	Id int
	Name     string `fako:"full_name"`
	Username string `fako:"user_name"`
	Email    string `fako:"email_address"`//Notice the fako:"email_address" tag
	Phone    string `fako:"phone"`
	Password string `fako:"simple_password"`
	Address  string `fako:"street_address"`
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
