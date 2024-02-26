package structs

import "fmt"

type user struct {
	id       int
	username string
	email    string
}

func User(username string, email string) user {
	return user{id: 1, username: username, email: email}
}

func (u user) GetId() int {
	return u.id
}

// Stringer interface implementation
func (u user) String() string {
	return fmt.Sprintf("%v. %v (%v)", u.id, u.username, u.email)
}
