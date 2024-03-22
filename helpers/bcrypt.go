package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Println("HashPass error:", err)
		return ""
	}

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)
	return err == nil
}
