package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateSalt() []byte {
	return []byte("beryodium")
}

func HashAndSaltPassword(password string, salt []byte) string {
	passwordBytes := []byte(password)
	saltedPassword := append(passwordBytes, salt...)
	sha := sha1.New()
	sha.Write(saltedPassword)
	saltedAndHashedPassword := sha.Sum(nil)
	return fmt.Sprintf("%x", saltedAndHashedPassword)
}

func PasswordMatch(hashedAndSaltedPassword string, password string) bool {
	return hashedAndSaltedPassword == HashAndSaltPassword(password, GenerateSalt())
}

type Claims struct {
	PhoneNumber string `json:"phoneNumber"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(phoneNumber string) (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	claims := &Claims{
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
