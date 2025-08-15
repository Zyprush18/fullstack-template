package helper

import "golang.org/x/crypto/bcrypt"

func HashingPass(pass string) ([]byte, error)  {
	return bcrypt.GenerateFromPassword([]byte(pass), 12)
}

func DecryptPass(hashpass,passreq string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpass),[]byte(passreq))
}