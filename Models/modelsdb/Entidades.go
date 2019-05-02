package modelsdb

import "time"

type Policy struct {
	Name string
	GUID string
	ColletionRefer  string     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column

}

type User struct {
	ID        string `gorm:"primary_key;uuid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time

	Politicas []Policy `gorm:"foreignkey:UserRefer"`

}