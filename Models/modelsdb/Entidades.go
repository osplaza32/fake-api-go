package modelsdb

import "time"

type Swagger struct{
	ID        string `gorm:"primary_key;uuid"`
	Content string
	Version string
	AuthorId string
	ServiceId string
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`



}

type Service struct {
	ID        string `gorm:"primary_key;uuid"`

	Name string
	BasePath string
	GUID string
	Content string
	Description string
	Swaggers []Swagger
	Policys []Policy
	UserRefer  string
	TimeStamp string
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`

}
type Policy struct {
	ID        string `gorm:"primary_key;uuid"`
	AuthorId string
	ServiceId string

	Name string
	GUID string
	Description string
	Version string
	NewVersion string
	Xml string
	TimeStamp string
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`

}


type User struct {
	ID        string `gorm:"primary_key;uuid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password"`
	Politicas []Policy
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`

}