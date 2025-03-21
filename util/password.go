package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",fmt.Errorf("faield to hash password: %w",err)
	}
	return string(hashedPassword),nil
}

func CheckPassword(pass string, hashedPass string)error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPass),[]byte(pass))
}