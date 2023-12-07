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
	RoleId      uint   `json:"role_id"`
	Role        Role   `json:"role" gorm:"foreignKey:RoleId"`
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

const (
	minFirstName = 5
	minLastName  = 5
	minPass      = 5
)

func (u *NewUser) SetPass(newPass string) {
	u.Password = newPass
}

func (u NewUser) ValidateLen() map[string]string {
	errors := map[string]string{}
	if len(u.FirstName) < minFirstName {
		errors["firstName"] = "first name should be 5 characteres length"
	}
	if len(u.LastName) < minLastName {
		errors["lastName"] = "last name should be 5 characteres length"
	}
	if len(u.Password) < minPass {
		errors["password"] = "last name should be 5 characteres length"
	}
	return errors

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
