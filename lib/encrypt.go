package lib

import "github.com/matthewhartstonge/argon2"

var Argon2 = argon2.DefaultConfig()

func Encrypt(password string) string {
	encoded, _ := Argon2.HashEncoded([]byte(password))

	return string(encoded)
}

func Verify(plainPassword string, encryptedPassword string) bool {
	ok, err := argon2.VerifyEncoded([]byte(plainPassword), []byte(encryptedPassword))
	if err != nil {
		return false
	}
	if ok {
		return true
	}
	return false
}
