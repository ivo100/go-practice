package userservice

// User represents a user in our service
type User struct {
	ID        string `json:"id" binding:"required,min=2"`
	FirstName string `json:"first_name" binding:"required,min=2"`
	LastName  string `json:"last_name" binding:"required,min=2"`
}

// UserService is the public interface of our service
type UserService interface {
	GetUser(id string) (*User, error)
	AddUser(user User) (*User, error)
}
