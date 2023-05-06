package core

// UserRepository is defined by service layer and implemented by repositories.
// It follows DIP principle (high level abstractions should not depend on low level abstractions).
type UserRepository interface {
	GetUser(id string) (*UserDto, error)
	SaveUser(user UserDto) (*UserDto, error)
}
