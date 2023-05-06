package repository

import (
	"myservice/internal/core"
	"sync"
)

// MemoryUserRepository is an in-memory implementation of UserRepository.
// it uses sync.Map so it is safe to use it concurrently.
type MemoryUserRepository struct {
	users sync.Map
}

func NewMemoryUserRepository() core.UserRepository {
	return &MemoryUserRepository{}
}

// GetUser retrieves existing user
// it returns nil if not found
func (r *MemoryUserRepository) GetUser(id string) (*core.UserDto, error) {
	user, ok := r.users.Load(id)
	if !ok {
		return nil, core.ErrNotFound
	}
	u, ok := user.(core.UserDto)
	if !ok {
		return nil, core.ErrInternal
	}
	return &u, nil

}

// SaveUser creates new user or updates existing user
func (r *MemoryUserRepository) SaveUser(user core.UserDto) error {
	r.users.Store(user.ID, user)
	return nil
}
