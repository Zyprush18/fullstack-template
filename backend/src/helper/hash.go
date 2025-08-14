package helper

import "golang.org/x/crypto/bcrypt"

func HashingPass(pass string) ([]byte, error)  {
	return bcrypt.GenerateFromPassword([]byte(pass), 12)
}