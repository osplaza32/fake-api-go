package Models

import "github.com/wawandco/fako"

var IdentityKey = "id"
type Caca struct {
	stuff interface{} // <- El camino a la gloria con las otras ideas!!!
}
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type UserFake struct {
	Id int
	Name     string `fako:"full_name"`
	Username string `fako:"user_name"`
	Email    string `fako:"email_address"`//Notice the fako:"email_address" tag
	Phone    string `fako:"phone"`
	Password string `fako:"simple_password"`
	Address  string `fako:"street_address"`
	Extras DataUser
}

type UserEdit struct {
	Name     string `json:"full_name,omitempty"`
	Username string `json:"user_name,omitempty"`
	Email    string `json:"email_address,omitempty"`//Notice the fako:"email_address" tag
	Phone    string `json:"phone,omitempty"`
	Password string `json:"simple_password,omitempty"`
	Address  string `json:"street_address,omitempty"`
	Extras DataUser	`json:"Extra,omitempty"`
}

type DataUser struct {
	Id int `json:"Id,omitempty"`
	Edad     int `json:"Edad,omitempty"`
	Compania string `json:"Company,omitempty"`
	ColorOjos    string `json:"color-eyes,omitempty"`
	ColorPelo    string `json:"color-hair,omitempty"`

}
type Users struct {
	App string
	Users []UserFake
}
func (u Users) Makeuser() Users {
	for i := 0; i < 10; i++ {
		u.Users = append(u.Users, makeUser())
		u.Users[i].Id = i
	}
	return u
}
func (u Users) CheckparameterFill(i int)bool {
	if  len(u.Users)-1 >= i {
		return true
	}
	return false
}
func makeUser() UserFake {
	var user UserFake
	fako.Fill(&user)
	return user

}