package main

import (
	"fmt"
	"osplaza32/ApiFastToTest/config"
)

func main() {
	db,err := Config.Conneccion()
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
	}
	fmt.Println("coneccion %+v",db)
}
