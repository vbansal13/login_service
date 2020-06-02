package service

import (
	"net/http"
	"testing"

	"github.com/vbansal/login_service/model"
)

func TestUserService_SignupUser(t *testing.T) {
	type args struct {
		user *model.UserRegisterRequestModel
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Missing username test case",
			args: args{
				&model.UserRegisterRequestModel{
					Password:  "Test123",
					Email:     "test@gmail.com",
					FirstName: "TestVarun",
				},
			},
			want:    http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Missing password test case",
			args: args{
				&model.UserRegisterRequestModel{
					Username:  "TestVB",
					Email:     "test@gmail.com",
					FirstName: "TestVarun",
				},
			},
			want:    http.StatusBadRequest,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uService := &UserService{}
			_, got1, err := uService.SignupUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.SignupUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want {
				t.Errorf("UserService.SignupUser() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}

func TestUserService_LoginUser(t *testing.T) {
	type args struct {
		user *model.UserRegisterRequestModel
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Invalid credentials",
			args: args{
				&model.UserRegisterRequestModel{
					Username: "vbansal1",
					Password: "Test123",
				},
			},
			want:    http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Valid credentials",
			args: args{
				&model.UserRegisterRequestModel{
					Username: "vbansal1",
					Password: "varun123",
				},
			},
			want:    http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uService := &UserService{}
			_, got1, err := uService.LoginUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.LoginUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want {
				t.Errorf("UserService.LoginUser() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}

func TestUserService_GetUserProfile(t *testing.T) {
	type args struct {
		accessTokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Invalid token",
			args: args{
				"dfdsfsdfadsfdasf",
			},
			want:    http.StatusUnauthorized,
			wantErr: true,
		},
		{
			name: "Valid token",
			args: args{
				"GciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRlIjoiMjAyMC0wMi0xNlQwNjo1NjoyNi0wODowMCIsInVzZXJuYW1lIjoidmJhbnNhbDQifQ.A4naOxK_FntuEmM1NZyyacfigy33_1ULksQ6ShEXRK8",
			},
			want:    http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uService := &UserService{}
			_, got1, err := uService.GetUserProfile(tt.args.accessTokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want {
				t.Errorf("UserService.GetUserProfile() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}
