package helpers

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(rawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hashPassword, unHashPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(unHashPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}
