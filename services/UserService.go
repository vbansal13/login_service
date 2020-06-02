package service

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	adapter "github.com/vbansal/login_service/adapter"
	"github.com/vbansal/login_service/analytics"
	"github.com/vbansal/login_service/config"
	"github.com/vbansal/login_service/dao"
	"github.com/vbansal/login_service/model"
	"golang.org/x/crypto/bcrypt"
)

//UserService data type
type UserService struct {
}

//NewUserService factory method for creating UserService object
func NewUserService() *UserService {
	return &UserService{}
}

//SignupUser registers a new user
//If successful it will return response model, if unsuccessful it will return contextual error message and http status code.
func (uService *UserService) SignupUser(user *model.UserRegisterRequestModel) (*model.ResponseResultModel, int, error) {

	err := uService.isMissingCriticalUserData(user, true)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	userDAO := dao.GetUserDAOInstance()

	_, err = userDAO.FindUser(user.Username)

	var response model.ResponseResultModel

	appConfig := config.GetInstance()

	if err != nil {
		if userDAO.UserNotFoundError(err) {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				return nil, http.StatusInternalServerError, fmt.Errorf(appConfig.Constants.HashingPasswordError)
			}
			user.Password = string(hash)
			tUser := adapter.NewUserAdapter().ConvertRequestModelToUserModel(user)
			tUser.CreationDate = time.Now()
			err = userDAO.AddNewUser(tUser)
			if err != nil {
				return nil, http.StatusInternalServerError, fmt.Errorf(appConfig.Constants.UserCreationError)
			}
			response.Result = appConfig.Constants.RegistrationSuccessful
			return &response, http.StatusOK, nil
		}
		return nil, http.StatusInternalServerError, err
	}
	response.Result = appConfig.Constants.UsernameExistsError
	return &response, http.StatusOK, nil
}

//LoginUser login a user and returns access token if successul
func (uService *UserService) LoginUser(user *model.UserRegisterRequestModel) (*model.UserLoginResponseModel, int, error) {
	//Check if request has all required user data
	err := uService.isMissingCriticalUserData(user, false)

	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	userDAO := dao.GetUserDAOInstance()
	userAnalytics := analytics.NewAnalytics()
	tUser, err := userDAO.FindUser(user.Username)
	appConfig := config.GetInstance()

	//If cannot find username than return Invalid credentials, to avoid username search attack.
	if err != nil {
		userAnalytics.UserLoginAttempt(false, user.Username, appConfig.Constants.InvalidUsername)
		return nil, http.StatusBadRequest, fmt.Errorf(appConfig.Constants.InvalidCredentialsError)
	}

	//If account is locked return error
	if tUser.AccountLocked {
		userAnalytics.UserLoginAttempt(false, user.Username, appConfig.Constants.AccountLockedError)
		return nil, http.StatusForbidden, fmt.Errorf(appConfig.Constants.AccountLockedError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(tUser.Password), []byte(user.Password))

	//This will be true due to password mismatch
	if err != nil {
		uService.saveLoginAttempt(false, tUser)
		userAnalytics.UserLoginAttempt(false, user.Username, appConfig.Constants.InvalidPassword)
		return nil, http.StatusBadRequest, fmt.Errorf(appConfig.Constants.InvalidCredentialsError)
	}

	accessTokenString, err := NewTokenService().GenerateSignedAccessToken(tUser.Username)

	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(appConfig.Constants.TokenGenerationError)
	}
	tUser.AccessToken = accessTokenString
	uService.saveLoginAttempt(true, tUser)

	userAnalytics.UserLoginAttempt(true, user.Username, appConfig.Constants.SuccessfulLogin)

	var result model.UserLoginResponseModel
	result.AccessToken = accessTokenString
	return &result, http.StatusOK, nil
}

//GetUserProfile fetches user profile, will return error if access token is valid or user is not found
func (uService *UserService) GetUserProfile(accessTokenString string) (*model.UserResponseModel, int, error) {
	appConfig := config.GetInstance()

	if len(accessTokenString) == 0 {
		return nil, http.StatusUnauthorized, fmt.Errorf(appConfig.Constants.MissingAuthTokenError)
	}

	username, err := NewTokenService().ValidateAccessTokenAndGetUser(accessTokenString)
	if err != nil {
		return nil, http.StatusUnauthorized, err
	}

	userDAO := dao.GetUserDAOInstance()
	tUser, uError := userDAO.FindUser(username)
	if uError != nil {
		return nil, http.StatusBadRequest, uError
	}

	return adapter.NewUserAdapter().ConvertUserModelToUserResponseModel(tUser), http.StatusOK, nil
}

func (uService *UserService) isMissingCriticalUserData(tUser *model.UserRegisterRequestModel, forSignup bool) error {
	appConfig := config.GetInstance()

	if tUser.Username == "" {
		return fmt.Errorf(appConfig.Constants.MissingUsernameError)
	}

	if tUser.Password == "" {
		return fmt.Errorf(appConfig.Constants.MissingPasswordError)
	}

	if !forSignup {
		return nil
	}

	if tUser.Email == "" {
		return fmt.Errorf(appConfig.Constants.MissingEmailError)
	}
	if tUser.FirstName == "" {
		return fmt.Errorf(appConfig.Constants.MissingFirstnameError)
	}
	return nil
}

//Internal helper method for saving successful and failed login attempts.
//Multiple consecutive failed login attempts will lock user's account.
func (uService *UserService) saveLoginAttempt(successful bool, user *model.UserModel) {

	if successful {
		user.UnsuccessfulLoginAttempts = 0
		user.LastLoginDate = time.Now()
	} else {
		appConfig := config.GetInstance()
		user.UnsuccessfulLoginAttempts++
		if user.UnsuccessfulLoginAttempts >= appConfig.MaxUnsuccessfulLoginAttemptsAllowed {
			user.AccountLocked = true
		}
	}
	userDAO := dao.GetUserDAOInstance()
	err := userDAO.SaveUser(user)
	if err != nil {
		log.Warning(err.Error())
	}
}
