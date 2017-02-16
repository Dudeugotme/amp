package schema

import (
	"fmt"
	"net/mail"
	"strings"
)

func isEmpty(s string) bool {
	return s == "" || strings.TrimSpace(s) == ""
}

func checkName(userName string) error {
	if isEmpty(userName) {
		return fmt.Errorf("name is mandatory")
	}
	return nil
}

func checkPasswordHash(passwordHash string) error {
	if isEmpty(passwordHash) {
		return fmt.Errorf("password hash is mandatory")
	}
	return nil
}

func checkEmail(email string) (string, error) {
	address, err := mail.ParseAddress(email)
	if err != nil {
		return "", fmt.Errorf("invalid email: %s", err.Error())
	}
	if isEmpty(address.Address) {
		return "", fmt.Errorf("email is mandatory")
	}
	return address.Address, nil
}

// Validate validates User
func (u *User) Validate() (err error) {
	if err = checkName(u.Name); err != nil {
		return err
	}
	if u.Email, err = checkEmail(u.Email); err != nil {
		return err
	}
	if err = checkPasswordHash(u.PasswordHash); err != nil {
		return err
	}
	return nil
}