package model

import "errors"

type User struct {
	ID   int
	Name string
	Role string
}

type Users []User

func (u Users) Exists(id int) bool {
	exists := false
	for _, user := range u {
		if user.ID == id {
			return true
		}
	}
	return exists
}

func (u Users) FindByName(name string) (User, error) {
	for _, user := range u {
		if user.Name == name {
			return user, nil
		}
	}
	return User{}, errors.New("USER_NOT_FOUND")
}

func CreateUsers() Users {
	users := Users{}
	users = append(users, User{ID: 1, Name: "Admin", Role: "admin"})
	users = append(users, User{ID: 2, Name: "Sabine", Role: "member"})
	users = append(users, User{ID: 3, Name: "Sepp", Role: "member"})
	return users
}
