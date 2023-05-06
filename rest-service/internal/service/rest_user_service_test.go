package service

import (
	"reflect"
	"testing"
	svc "userservice/pkg"
)

func TestUserServiceImpl_AddGetUser(t *testing.T) {
	user := svc.User{
		ID:        "1",
		FirstName: "a",
		LastName:  "b",
	}
	s := NewUserService()
	tests := []struct {
		name    string
		want    *svc.User
		wantErr bool
	}{
		{
			name:    "add user",
			want:    &user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AddUser(user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUser() got = %v, want %v", got, tt.want)
			}

			got, err = s.GetUser(user.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}

		})
	}
}
