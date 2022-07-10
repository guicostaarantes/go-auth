package hash_util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type BCryptImpl struct {
	Cost int
}

func (u BCryptImpl) Hash(plain string) (string, error) {
	hashedBytes, hashErr := bcrypt.GenerateFromPassword([]byte(plain), u.Cost)
	if hashErr != nil {
		return "", errors.New("internal server error")
	}
	return string(hashedBytes), nil
}

func (u BCryptImpl) Compare(plain string, hashed string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
