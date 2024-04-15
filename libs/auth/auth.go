package auth

import (
	"github.com/sandisuryadi36/sansan-dashboard/libs/logger"
	"golang.org/x/crypto/bcrypt"
)

// function to generate hashed password from string password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Error hashing password: ", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// function to compare string password and hashed password
func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
