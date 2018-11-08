package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func CheckPasswordHash(passwordHashed, password *string) (error, bool) {
	err := bcrypt.CompareHashAndPassword([]byte(*passwordHashed), []byte(*password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, false
		}
		return err, false
	}
	return nil, true
}
