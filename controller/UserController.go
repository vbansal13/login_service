package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vbansal/login_service/model"
	service "github.com/vbansal/login_service/services"
)

//UserController structure
type UserController struct {
}

//NewUserController factory method for creating UserController object
func NewUserController() *UserController {
	return &UserController{}
}

// SignupUserHandler godoc
// @Summary Registers a new user
// @Produce json
// @Param request body model.UserRegisterRequestModel true "user info"
// @Success 200 {object} model.ResponseResultModel
// @Failure 400 {object}  model.ResponseResultModel
// @Router /me [post]
//SignupUserHandler Method for handling sign-up request and creating new user account in DB
func (ctr *UserController) SignupUserHandler(c *gin.Context) {

	user, err := ctr.getUserObjectFromRequest(c)

	if err != nil {
		encodeErrorResponse(err.Error(), c, http.StatusBadRequest)
		return
	}

	response, statusCode, sErr := service.NewUserService().SignupUser(user)

	if sErr != nil {
		encodeErrorResponse(sErr.Error(), c, statusCode)
		return
	}
	c.JSON(http.StatusOK, response)
}

// LoginUserHandler godoc
// @Summary Login a user
// @Produce json
// @Param request body model.UserLoginRequestModel true "user credentials"
// @Success 200 {object} model.UserLoginResponseModel
// @Failure 400 {object}  model.ResponseResultModel
// @Failure 403 {object}  model.ResponseResultModel
// @Failure 403 {object}  model.ResponseResultModel
// @Router /me/login [post]
//LoginUserHandler Method for handling login requests
func (ctr *UserController) LoginUserHandler(c *gin.Context) {

	user, err := ctr.getUserObjectFromRequest(c)
	if err != nil {
		encodeErrorResponse(err.Error(), c, http.StatusBadRequest)
		return
	}
	response, statusCode, sErr := service.NewUserService().LoginUser(user)

	if sErr != nil {
		encodeErrorResponse(sErr.Error(), c, statusCode)
		return
	}

	c.JSON(http.StatusOK, response)
}

// ProfileUserHandler godoc
// @Summary Fetch user profile
// @Produce json
// @Success 200 {object} model.UserResponseModel
// @Failure 401 {object}  model.ResponseResultModel
// @Router /me [get]
//ProfileUserHandler Method for handling fetch user profile request
func (ctr *UserController) ProfileUserHandler(c *gin.Context) {

	accessTokenString := c.GetHeader("Authorization")
	user, statusCode, err := service.NewUserService().GetUserProfile(accessTokenString)

	if err != nil {
		encodeErrorResponse(err.Error(), c, statusCode)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctr *UserController) getUserObjectFromRequest(c *gin.Context) (*model.UserRegisterRequestModel, error) {

	var user model.UserRegisterRequestModel
	err := c.ShouldBindJSON(&user)
	return &user, err
}

func encodeErrorResponse(errString string, c *gin.Context, statusCode int) {
	var res model.ResponseResultModel
	res.Error = errString

	c.JSON(statusCode, res)
}
