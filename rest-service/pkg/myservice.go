package myservice

import (
	_ "github.com/go-playground/validator/v10"
)

// User represents a user in our service
type User struct {
	ID        string `json:"id", validate:"required"`
	FirstName string `json:"first_name", validate:"required"`
	LastName  string `json:"last_name", validate:"required"`
}

// MyService is the public interface of our service exposed via REST
type MyService interface {
	GetUser(id string) (*User, error)
	AddUser(user User) error
}
