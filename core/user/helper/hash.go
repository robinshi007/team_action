package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), plainPwd)
	if err != nil {
		return false
	}
	return true
}
