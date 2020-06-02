package service

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vbansal/login_service/config"
)

//TokenService data type
type TokenService struct {
}

//NewTokenService factory method for creating token service object
func NewTokenService() *TokenService {
	return &TokenService{}
}

//GenerateSignedAccessToken method for generating signed access token.
func (ctr *TokenService) GenerateSignedAccessToken(username string) (string, error) {
	tTime := time.Now()
	appConfig := config.GetInstance()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"date":     tTime.Format(time.RFC3339),
	})

	return token.SignedString([]byte(appConfig.SigningSecret))
}

//ValidateAccessTokenAndGetUser method of validating access token and fetch username if token valid.
func (ctr *TokenService) ValidateAccessTokenAndGetUser(accessTokenString string) (string, error) {

	appConfig := config.GetInstance()

	accessToken, err := jwt.Parse(accessTokenString, func(accessToken *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := accessToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(appConfig.Constants.SigningError)
		}
		return []byte(appConfig.SigningSecret), nil
	})

	if err != nil {
		return "", err
	}
	if claims, ok := accessToken.Claims.(jwt.MapClaims); ok && accessToken.Valid {
		tokenDateString := claims["date"].(string)
		currentTime := time.Now()
		tokenTime, tErr := time.Parse(time.RFC3339, tokenDateString)

		if tErr != nil {
			return "", fmt.Errorf(appConfig.Constants.InvalidTokenError)
		}

		//The token is only good for 1hr from generate date/
		expireAt := tokenTime.Add(3600 * time.Second)
		if currentTime.After(expireAt) {
			return "", fmt.Errorf(appConfig.Constants.InvalidTokenError)
		}

		username := claims["username"].(string)
		return username, nil
	}
	return "", err

}
