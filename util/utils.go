package util

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func HashAndSaltPassword(password string) string {
	passwordBytes := []byte(password)
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltBytes := []byte(salt)
	saltedPassword := append(passwordBytes, saltBytes...)
	sha := sha1.New()
	sha.Write(saltedPassword)
	saltedAndHashedPassword := sha.Sum(nil)
	return fmt.Sprintf("%x", saltedAndHashedPassword)
}
