package types

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

type User struct {
	Id          uint   `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" gorm:"unique"`
	EncPassword string `json:"encPassword"`
}

type NewUser struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
}

type UserLoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func NewUserFn(newUser NewUser) (*User, error) {
	encPw, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	u := User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		EncPassword: string(encPw),
	}

	return &u, nil
}
