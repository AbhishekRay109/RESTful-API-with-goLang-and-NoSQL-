package main

type UserService interface {
	CreateUser(*User) error
	GetUser(*string) (*User, error)
	GetAll() ([]*User, error)
	UpdateUser(*User) error
	DeleteUser(*string) error
}

// List of all the Actions(Functions)
