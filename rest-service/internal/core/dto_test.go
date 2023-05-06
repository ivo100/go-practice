package core

import (
	"reflect"
	"testing"
)

func TestMapToUserDto(t *testing.T) {
	type args struct {
		input map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    *UserDto
		wantErr bool
	}{
		{
			name: "first",
			args: args{
				input: map[string]any{"id": "1", "first_name": "a", "last_name": "b"},
			},
			want: &UserDto{
				ID:        "1",
				FirstName: "a",
				LastName:  "b",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapToUserDto(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToUserDto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToUserDto() got = %v, want %v", got, tt.want)
			}
		})
	}
}
