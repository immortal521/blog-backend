package entity

import (
	"fmt"

	"blog-server/ent"
	"blog-server/utils"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ent.User
}

func NewUser(email, password string) (*User, error) {
	u := &User{
		ent.User{
			Email: email,
		},
	}

	pwd, err := u.HashPassword(password)
	if err != nil {
		return nil, err
	}

	u.Password = pwd
	u.Username = u.GenerateUsername()

	return u, nil
}

func (u *User) HashPassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 12)
	if err != nil {
		return "", fmt.Errorf("hash password failed: %w", err)
	}
	return string(hash), nil
}

func ValidatePassword(u *User, p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return nil
	}
	return nil
}

func (u *User) GenerateUsername() string {
	return utils.RandomString(10, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}
