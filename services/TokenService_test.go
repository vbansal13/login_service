package service

import (
	"testing"
)

func TestTokenService_GenerateSignedAccessToken(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		ctr     *TokenService
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GenerateSignedToken",
			ctr:  nil,
			args: args{
				"vbansal4",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := &TokenService{}
			_, err := ctr.GenerateSignedAccessToken(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.generateSignedAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTokenService_ValidateAccessTokenAndGetUser(t *testing.T) {
	type args struct {
		accessTokenString string
	}
	tests := []struct {
		name    string
		ctr     *TokenService
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ValidToken",
			ctr:  nil,
			args: args{
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRlIjoiMjAyMC0wMi0xNlQwNjo1NjoyNi0wODowMCIsInVzZXJuYW1lIjoidmJhbnNhbDQifQ.A4naOxK_FntuEmM1NZyyacfigy33_1ULksQ6ShEXRK8",
			},
			want:    "vbansal4",
			wantErr: false,
		},
		{
			name: "InValidToken",
			ctr:  nil,
			args: args{
				"GciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRlIjoiMjAyMC0wMi0xNlQwNjo1NjoyNi0wODowMCIsInVzZXJuYW1lIjoidmJhbnNhbDQifQ.A4naOxK_FntuEmM1NZyyacfigy33_1ULksQ6ShEXRK8",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := &TokenService{}
			got, err := ctr.ValidateAccessTokenAndGetUser(tt.args.accessTokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.validateAccessTokenAndGetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TokenService.validateAccessTokenAndGetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
