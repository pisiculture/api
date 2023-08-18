package domain

import "errors"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func NewUser() *User {
	return &User{}
}

func (usr *User) Validate() error {
	if usr.Name == "" {
		return errors.New("Param name is not valid.")
	}
	if usr.Email == "" {
		return errors.New("Param e-mail is not valid.")
	}
	if usr.Password == "" || len(usr.Password) < 8 {
		return errors.New("Param password is not valid.")
	}
	return nil
}
