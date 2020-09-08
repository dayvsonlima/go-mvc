package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User .
type User struct {
	gorm.Model `json:"-"`
	Email      string `gorm:"unique;index"`
	Password   string `json:"-" binding:"required"`
}

// SetPassword Set user password of the right way
func (u *User) SetPassword(password string) {
	secret := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(secret, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.Password = string(passwordHash)
}

// SetEmail Set user password of the right way
func (u *User) SetEmail(email string) {
	u.Email = email
}

// CheckPassword .
// check if password correspond to inputed
func (u *User) CheckPassword(password string) bool {
	secret := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(secret, []byte(password))

	return err == nil
}
