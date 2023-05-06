package repository

import (
	"reflect"
	"testing"
	"userservice/internal/core"
)

func TestMemoryUserRepository_SaveAndGetUser(t *testing.T) {

	user := core.UserDto{
		ID:        "1",
		FirstName: "a",
		LastName:  "b",
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "first",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMemoryUserRepository()
			got, err := r.SaveUser(user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, user) {
				t.Errorf("GetUser() got = %v, want %v", got, user)
			}
			got, err = r.GetUser(user.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, user) {
				t.Errorf("GetUser() got = %v, want %v", got, user)
			}
		})
	}
}
