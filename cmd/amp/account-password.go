package main

import (
	"fmt"
	"github.com/appcelerator/amp/api/client"
	backend "github.com/appcelerator/amp/api/rpc/account"
	"github.com/appcelerator/amp/cmd/amp/account"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

const resetVal = "reset"
const changeVal = "change"

var (
	AccountCmd = &cobra.Command{
		Use:   "account",
		Short: "Account operations",
		Long:  `The account command manages all account-related operations.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return AMP.Connect()
		},
	}

	pwdCmd = &cobra.Command{
		Use:   "password",
		Short: "Password operations",
		Long:  `The reset command resets the password of an account.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return pwd(AMP, cmd, args)
		},
	}

	reset      bool
	change     bool
	err        error
	cmdSpecArr []account.CommandSpec

	dir = "./cmd/amp/account/commands"
)

func init() {
	cmdSpecArr, err = account.ParseFile()
	if err != nil {
		fmt.Println("parsing error : ", err)
	}
	RootCmd.AddCommand(AccountCmd)
	AccountCmd.AddCommand(pwdCmd)
	pwdCmd.Flags().BoolVar(&reset, resetVal, false, "Reset Password")
	pwdCmd.Flags().BoolVar(&change, changeVal, false, "Change Password")
}

func pwd(amp *client.AMP, cmd *cobra.Command, args []string) (err error) {
	var argArr []string
	var username, email string
	//Iterating over command spec
	for _, cmdSpec := range cmdSpecArr {
		//Iterating over flags
		for _, flag := range cmdSpec.Flags {
			//Iterating over reset flag info
			for _, detail := range flag.Reset {
				if detail.Value == resetVal && detail.Required {
					fmt.Println(detail.Info)
				}
				if reset {
					switch len(args) {
					case 0:
						username = account.GetUserName()
						email, err = account.GetEmailAddress()
						if err != nil {
							return fmt.Errorf("user error: %v", err)
						}
						argArr = append(argArr, username, email)
						err = account.ValidateArguments(argArr, cmdSpecArr)
						if err != nil {
							return fmt.Errorf("user error : %v", err)
						}

					case 1:
						username = args[0]
						email, err = account.GetEmailAddress()
						if err != nil {
							return fmt.Errorf("user error: %v", err)
						}
						argArr = append(argArr, username, email)
						err = account.ValidateArguments(argArr, cmdSpecArr)
						if err != nil {
							return fmt.Errorf("user error : %v", err)
						}

					case 2:
						argArr = append(argArr, args...)
						err = account.ValidateArguments(argArr, cmdSpecArr)
						if err != nil {
							return fmt.Errorf("user error : %v", err)
						}
					default:
						return fmt.Errorf("too many arguments")
					}
					request := &backend.PasswordResetRequest{
						Name:  username,
						Email: email,
					}
					client := backend.NewAccountClient(amp.Conn)
					_, err = client.PasswordReset(context.Background(), request)
					if err != nil {
						return fmt.Errorf("server error: %v", err)
					}
				} else if change {
					fmt.Println("pwd change to be implemented")
				} else {
					fmt.Println("Incomplete command - Choose a flag. Try again!")
				}
			}
		}
	}

	return nil
}
