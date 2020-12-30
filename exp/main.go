package main

import (
	"fmt"

	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	// us.DestructiveReset()
	if err != nil {
		panic(err)

	}
	// user := models.User{
	// 	Name:  "Zubin",
	// 	Email: "r.zubin@protonmail.com",
	// }
	// us.Create(&user)
	user, err := us.ByID(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
