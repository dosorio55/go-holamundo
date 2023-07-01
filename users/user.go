package users

import (
	"fmt"
	"github.com/dosorio55/hola-mundo/models"
	// "time"
)

func CreateUser() {
	// newUser := new(models.User)
	newUser := models.User{
		Id: 1,
		Name: "David",
		// CreatedAt: time.Now(),
		Status: true,
	}

	newUser.UpdateName("Karla")


	// newUser.AddUser(1, "David", time.Now(), true)

	fmt.Println(newUser)
}
