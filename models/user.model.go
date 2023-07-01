package models

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	// CreatedAt time.Time
	Status    bool
}

// la este es que antes de la función ponemos el nombre de la estructura a la que se está refiriendo
// la palabra this realmente puede ser cualquier cosa, pero por convención se usa una sola letra
// en este caso se usa this para poder relacionarlo con otros lenguajes de programación que si usan la palabra this
func(this *User) AddUser(id int, name string, createdAt time.Time, Status bool) {
	// the * is a pointer to the User struct
	// it means it chares the same memory address
	// like objects in js
	// without it, the values are not going to be changed in this same object
	this.Id = id
	this.Name = name
	// this.CreatedAt = createdAt
	this.Status = Status
}

func(u User) UpdateName(name string) {
	 u.Name = name
}
