package code

func HashPassword() string {
	return `
package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hash)
}

func ComparePassword(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
`
}
