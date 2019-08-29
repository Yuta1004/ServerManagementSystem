package authfunc

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// GenPasswordHash : パスワードのハッシュを生成して返す
func GenPasswordHash(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
	}
	return string(hash)
}

// AuthPassword : パスワードとハッシュ化されたパスワードが合致するかを返す
func AuthPassword(rawPassword, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))
	return err == nil
}