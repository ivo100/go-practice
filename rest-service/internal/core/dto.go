package core

import (
	//	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

// UserDto is user data transfer object - it is used to communicate between service and repository layers
type UserDto struct {
	ID        string `mapstructure:"id"`
	FirstName string `mapstructure:"first_name"`
	LastName  string `mapstructure:"last_name"`
}

func MapToUserDto(input map[string]any) (*UserDto, error) {
	var res UserDto
	err := mapstructure.Decode(input, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
