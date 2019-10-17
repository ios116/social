package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword for save to bd
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash comparison password
func CheckPasswordHash(password, hash string) bool {
	fmt.Println("======>",password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
