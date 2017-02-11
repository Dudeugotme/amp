package account

import (
	"net/mail"

	//"github.com/holys/safe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"strings"
)

func isEmpty(s string) bool {
	return s == "" || strings.TrimSpace(s) == ""
}

func checkUserName(userName string) error {
	if isEmpty(userName) {
		return grpc.Errorf(codes.InvalidArgument, "name is mandatory")
	}
	return nil
}

func checkPassword(password string) error {
	if isEmpty(password) {
		return grpc.Errorf(codes.InvalidArgument, "password is mandatory")
	}
	// TODO: check password strength
	//err = CheckPasswordStrength(r.Password)
	//if err != nil {
	//	return
	//}
	return nil
}

func checkOrganizationName(name string) error {
	if isEmpty(name) {
		return grpc.Errorf(codes.InvalidArgument, "organization name is mandatory")
	}
	return nil
}

func checkEmail(email string) (string, error) {
	address, err := mail.ParseAddress(email)
	if err != nil {
		return "", grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	if isEmpty(address.Address) {
		return "", grpc.Errorf(codes.InvalidArgument, "email is mandatory")
	}
	return address.Address, nil
}

//func CheckPasswordStrength(password string) error {
//	s := safe.New(8, 0, 0, safe.Simple)
//	err := s.SetWords("./vendor/github.com/holys/safe/words.dat")
//	if err != nil {
//		return grpc.Errorf(codes.NotFound, "cannot find common password data")
//	}
//	str := s.Check(password)
//	if str < 2 {
//		return grpc.Errorf(codes.InvalidArgument, "password too weak")
//	}
//	return nil
//}

// Validate validates SignUpRequest
func (r *SignUpRequest) Validate() (err error) {
	if r.Email, err = checkEmail(r.Email); err != nil {
		return err
	}
	if err = checkPassword(r.Password); err != nil {
		return err
	}
	if err = checkUserName(r.UserName); err != nil {
		return err
	}
	return nil
}

// Validate validates VerificationRequest
func (r *VerificationRequest) Validate() error {
	return nil
}

//func (r *OrganizationRequest) Validate() (err error) {
//	err = checkOrganizationName(r.Name)
//	if err != nil {
//		return
//	}
//	email, err := CheckEmailAddress(r.Email)
//	if err != nil {
//		return
//	}
//	r.Email = email
//	return
//}

// Validate validates LogInRequest
func (r *LogInRequest) Validate() (err error) {
	if err = checkUserName(r.UserName); err != nil {
		return err
	}
	if err = checkPassword(r.Password); err != nil {
		return err
	}
	return nil
}

// Validate validates EditRequest
func (r *EditAccountRequest) Validate() (err error) {
	if err = checkUserName(r.UserName); err != nil {
		return err
	}
	if r.Email != "" {
		var email string
		email, err = checkEmail(r.Email)
		if err != nil {
			return
		}
		r.Email = email
	}
	//if r.NewPassword != "" {
	//	err = CheckPasswordStrength(r.NewPassword)
	//	if err != nil {
	//		return
	//	}
	//}
	return
}

// Validate validates PasswordResetRequest
func (r *PasswordResetRequest) Validate() (err error) {
	if err = checkUserName(r.UserName); err != nil {
		return err
	}
	if r.Email, err = checkEmail(r.Email); err != nil {
		return err
	}
	return
}