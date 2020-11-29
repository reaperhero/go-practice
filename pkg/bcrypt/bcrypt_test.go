package bcrypt

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func Test_bcrypt_01(T *testing.T) {
	pwd := "pa55w0rd"

	logrus.Println("Origin Password: " + pwd)

	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	logrus.Println("Encrypted Password: " + string(hash))

	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	logrus.Println("Match Result: ", err == nil)
}

func Test_CompareHashAndPassword(t *testing.T) {
	pwd := "pa55w0rd"
	hash := "$2a$10$8/ETBRr1aMdA3jYdG4q02Ou8tE1a5lSmktej97pq2PrsSGofmqOJq"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	logrus.Info(err)
}
