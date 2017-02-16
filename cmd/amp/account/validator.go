package account

import (
	"errors"
	"fmt"
	"net/mail"
)

func checkUserName(username string) (err error) {
	if username == "" {
		return errors.New("username is mandatory")
	}
	return nil
}

func checkEmailAddress(email string) (processedEmail string, err error) {
	address, err := mail.ParseAddress(email)
	if err != nil {
		return "", errors.New("invalid email format")
	}
	return address.Address, nil
}

func checkArgumentLength(arrLen int, cmdSpecLen int) (err error) {
	if arrLen != cmdSpecLen {
		return errors.New("too many arguments")
	}
	return
}

func checkLength(password string) (err error) {
	if len(password) != 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return
}

func ValidateArguments(cmdArr []string, cmdSpecArr []CommandSpec) (err error) {
	//fmt.Println(cmdArr)
	for _, cmdSpec := range cmdSpecArr {
		for _, argSpec := range cmdSpec.Args {
			//fmt.Println(len(cmdArr), " === ", cmdSpec.Length)
			err = checkArgumentLength(len(cmdArr), cmdSpec.Length)
			if err != nil {
				return err
			}
			for _, arg := range argSpec.Accname {
				err = checkUserName(cmdArr[0])
				if err != nil {
					return err
				} else {
					fmt.Println(arg.Response)
				}
			}
			//fmt.Println(cmdArr[1])
			for _, arg := range argSpec.Email {
				_, err = checkEmailAddress(cmdArr[1])
				if err != nil {
					return err
				} else {
					fmt.Println(arg.Response)
				}
			}
		}

	}
	return
}
