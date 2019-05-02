package Models

import "github.com/wawandco/fako"

type Serviceinfo struct {
	Item struct {
		Link []struct {
			URI string `json:"-uri,omitempty"`
			Rel string `json:"-rel,omitempty"`

		} `json:"Link"`
		Resource struct {
			Service struct {
				ID            string `json:"-id"`
				Version       string `json:"-version"`
				ServiceDetail struct {
					ServiceMappings struct {
						HTTPMapping struct {
							URLPattern string `json:"UrlPattern"`
							Verbs      struct {
								Verb []string `json:"Verb"`
							} `json:"Verbs"`
						} `json:"HttpMapping"`
					} `json:"ServiceMappings"`
					Properties struct {
						Property []struct {
							Key          string `json:"-key"`
							BooleanValue string `json:"BooleanValue,omitempty"`
							LongValue    string `json:"LongValue,omitempty"`
						} `json:"Property"`
					} `json:"Properties"`
					FolderID string `json:"-folderId"`
					ID       string `json:"-id"`
					Version  string `json:"-version"`
					Name     string `json:"Name"`
					Enabled  string `json:"Enabled"`
				} `json:"ServiceDetail"`
				Resources struct {
					ResourceSet struct {
						Resource struct {
							Content string `json:"#content"`
							Type    string `json:"-type"`
							Version string `json:"-version"`
						} `json:"Resource"`
						Tag string `json:"-tag"`
					} `json:"ResourceSet"`
				} `json:"Resources"`
			} `json:"Service"`
		} `json:"Resource"`
		L7        string `json:"-l7"`
		Name      string `json:"Name"`
		ID        string `json:"Id"`
		Type      string `json:"Type"`
		TimeStamp string `json:"TimeStamp"`
	} `json:"Item"`
}


type Policyinfo struct {
	Item struct {
		Name      string `json:"Name"`
		ID        string `json:"Id"`
		Type      string `json:"Type"`
		TimeStamp string `json:"TimeStamp"`
		Link      []struct {
			Rel string `json:"-rel"`
			URI string `json:"-uri,omitempty"`
		} `json:"Link"`
		Resource struct {
			Policy struct {
				GUID         string `json:"-guid"`
				ID           string `json:"-id"`
				Version      string `json:"-version"`
				PolicyDetail struct {
					ID         string `json:"-id"`
					Version    string `json:"-version"`
					Name       string `json:"Name"`
					PolicyType string `json:"PolicyType"`
					Properties struct {
						Property []struct {
							Key          string `json:"-key"`
							LongValue    string `json:"LongValue,omitempty"`
							BooleanValue string `json:"BooleanValue,omitempty"`
						} `json:"Property"`
					} `json:"Properties"`
					FolderID string `json:"-folderId"`
					GUID     string `json:"-guid"`
				} `json:"PolicyDetail"`
				Resources struct {
					ResourceSet struct {
						Resource struct {
							Content string `json:"#content"`
							Type    string `json:"-type"`
						} `json:"Resource"`
						Tag string `json:"-tag"`
					} `json:"ResourceSet"`
				} `json:"Resources"`
			} `json:"Policy"`
		} `json:"Resource"`
		L7 string `json:"-l7"`
	} `json:"Item"`

}
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