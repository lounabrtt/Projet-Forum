package validators

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(password, confirmPassword string) bool {
	return password == confirmPassword
}

// HashPassword hache le mot de passe en utilisant bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compare un mot de passe en texte clair avec un hachage
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
