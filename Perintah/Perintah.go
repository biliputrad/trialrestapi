package Perintah

import (
	"RestAPI-GETNPOST/Entity"
	"RestAPI-GETNPOST/Repository"
	"fmt"
)

//CRUD Product

//CRUD User
func GetPasswordUser(repository Repository.UserRepository, pass string) bool {
	password , err := repository.FindPasswordUser(pass)

	if err != nil {
		panic("Failed to get cart")
	}

	if pass == password {
		return true
	}
	else {
		return false
	}
}
func GetUsernameUser(repository Repository.UserRepository, username string) bool {
	name, err := repository.FindUsernameUser(username)

	if err != nil {
		panic("Failed to get username")
	}
	if username == name {
		return true
	}
	else {
		return false
	}
}

//CRUD Shopping cart

//CRUD Cart
