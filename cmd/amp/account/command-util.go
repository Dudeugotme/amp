package account

import (
	"fmt"
)

func GetUserName() (username string) {
	fmt.Print("username: ")
	fmt.Scanln(&username)
	if username == "" {
		return GetUserName()
	}
	return
}

func GetEmailAddress() (email string, err error) {
	fmt.Print("email: ")
	fmt.Scanln(&email)
	if email == "" {
		return GetEmailAddress()
	}
	return
}
