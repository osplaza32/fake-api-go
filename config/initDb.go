package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/pborman/uuid"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"osplaza32/ApiFastToTest/Models/modelsdb"
)

func Init(){
	db,err := Conneccion()
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
	}
	fmt.Println("coneccion %+v",db)
	db.Exec("CREATE EXTENSION postgis")
	db.Exec("CREATE EXTENSION postgis_topology")
	db.DropTable(&modelsdb.User{})

	db.AutoMigrate(&modelsdb.User{})
	db.DropTable(&modelsdb.Swagger{})

	db.AutoMigrate(&modelsdb.Swagger{})
	db.DropTable(&modelsdb.Service{})

	db.AutoMigrate(&modelsdb.Service{})
	db.DropTable(&modelsdb.Policy{})

	db.AutoMigrate(&modelsdb.Policy{})
	seed(db)

}
func Conneccion() (*gorm.DB,error) {
	gotenv.Load()
	var as string
	as = "host="+os.Getenv("HOST")+" user="+os.Getenv("PG_USER")+" dbname="+os.Getenv("PG_DB")+" sslmode=disable password="+os.Getenv("PG_PASS")
	db, err := gorm.Open("postgres", as)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
	}
	db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", beforeCreate)

	db.LogMode(true)
	return db,err
}

func seed(db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456789"), bcrypt.DefaultCost)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))

	}
	user := &modelsdb.User{FirstName:"Oscar",LastName:"Plaza",Email:"oscar.plaza@techo.org",Password:string(hashedPassword[:])}
	db.Create(&user)



}
func beforeCreate(scope *gorm.Scope) {
	scope.SetColumn("id", uuid.NewUUID().String())

}