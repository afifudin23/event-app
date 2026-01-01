package security

import "github.com/alexedwards/argon2id"

func HashPassword(password string) (string, error) {
	hashedPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	return hashedPassword, err
}

func CheckPassword(password string, hashedPassword string) bool {
	if _, err := argon2id.ComparePasswordAndHash(password, hashedPassword); err != nil {
		return false
	}
	return true
}
