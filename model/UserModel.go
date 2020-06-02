package model

import "time"

//UserModel desribes user information stored in user DB
type UserModel struct {
	Username                  string    `json:"username,omitempty"`
	FirstName                 string    `json:"firstname,omitempty"`
	LastName                  string    `json:"lastname,omitempty"`
	Email                     string    `json:"email,omitempty"`
	Password                  string    `json:"password,omitempty"`
	AccessToken               string    `json:"access_token,omitempty"`
	LastLoginDate             time.Time `json:"-"`
	CreationDate              time.Time `json:"-"`
	AccountLocked             bool      `json:"-"`
	UnsuccessfulLoginAttempts uint8     `json:"-"`
}

//UserLoginRequestModel desribes model for user login
type UserLoginRequestModel struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

//UserRegisterRequestModel desribes model for user registration
type UserRegisterRequestModel struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

//UserResponseModel describes response model for Profile request
type UserResponseModel struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

//UserLoginResponseModel describes response model for Login request
type UserLoginResponseModel struct {
	AccessToken string `json:"access_token,omitempty"`
}
