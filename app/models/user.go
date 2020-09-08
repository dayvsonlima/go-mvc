package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User .
type User struct {
	gorm.Model `json:"-"`
	email      string `binding:"required"`
	password   string `binding:"required"`
}

// SetPassword Set user password of the right way
func (u User) SetPassword(password string) {
	secret := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(secret, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.password = string(passwordHash)
}

// SetEmail Set user password of the right way
func (u User) SetEmail(email string) {
	u.email = email
}

// CheckPassword .
// check if password correspond to inputed
func (u User) CheckPassword(password string) bool {
	secret := []byte(u.password)
	err := bcrypt.CompareHashAndPassword(secret, []byte(password))

	return err == nil
}
