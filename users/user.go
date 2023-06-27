package users

import (
	"fmt"
	"github.com/dosorio55/hola-mundo/models"
	"time"
)

func CreateUser() {
	newUser := new(models.User)

	newUser.AddUser(1, "David", time.Now(), true)

	fmt.Println(newUser)
}
