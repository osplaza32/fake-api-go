package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"os"
	"osplaza32/ApiFastToTest/Models/modelsdb"
	"reflect"
	"strings"
)

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
	db.Exec("CREATE EXTENSION postgis")
	db.Exec("CREATE EXTENSION postgis_topology")
	db.AutoMigrate(&modelsdb.User{})
	db.AutoMigrate(&modelsdb.Policy{})


	return db,err
}
func beforeCreate(scope *gorm.Scope) {
	reflectValue := reflect.Indirect(reflect.ValueOf(scope.Value))
	if strings.Contains(string(reflectValue.Type().Field(0).Tag), "uuid") {
		uuid.SetClockSequence(-1)
		scope.SetColumn("id", uuid.NewUUID().String())
	}
}