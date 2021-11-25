package validation

import (
	pb "github.com/ash0tych/gRPC_1/proto/compiled/authorization"
	"net/mail"
)

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validatePassword(pass string) bool {
	return len([]rune(pass)) < 8
}

func CorrectInput(in *pb.UserData) (bool, *pb.UserError) {
	IsValidEmail := validateEmail(in.Login)
	if !IsValidEmail {
		return false, &pb.UserError{
			Err: "Incorrect Email",
			Id:  1,
		}
	}

	IsValidPassword := validatePassword(in.Password)
	if !IsValidPassword {
		return false, &pb.UserError{
			Err: "Incorrect Password",
			Id:  2,
		}
	}
	return true, nil
}
