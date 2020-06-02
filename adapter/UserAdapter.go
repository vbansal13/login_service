package adapter

import "github.com/vbansal/login_service/model"

//UserAdapter data type
type UserAdapter struct {
}

//NewUserAdapter factory method for creating user adapator object
func NewUserAdapter() *UserAdapter {
	return &UserAdapter{}
}

//ConvertRequestModelToUserModel converts registration request model to User Model
func (adp *UserAdapter) ConvertRequestModelToUserModel(reqModel *model.UserRegisterRequestModel) *model.UserModel {
	tUser := &model.UserModel{
		Username:  reqModel.Username,
		Password:  reqModel.Password,
		Email:     reqModel.Email,
		FirstName: reqModel.FirstName,
		LastName:  reqModel.FirstName,
	}
	return tUser
}

//ConvertUserModelToUserResponseModel converts User Model to User response model
func (adp *UserAdapter) ConvertUserModelToUserResponseModel(user *model.UserModel) *model.UserResponseModel {
	tUser := &model.UserResponseModel{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return tUser
}
