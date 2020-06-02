package config

//AppConstants defines string constants used accross the application
type AppConstants struct {
	HashingPasswordError    string
	UserCreationError       string
	RegistrationSuccessful  string
	UsernameExistsError     string
	InvalidCredentialsError string
	InvalidCredential       string
	InvalidUsername         string
	InvalidPassword         string
	AccountLockedError      string
	TokenGenerationError    string
	SuccessfulLogin         string
	MissingAuthTokenError   string
	MissingUsernameError    string
	MissingPasswordError    string
	MissingEmailError       string
	MissingFirstnameError   string
	InvalidTokenError       string
	SigningError            string
}

//Config defines Application configuration
type Config struct {
	Constants                           *AppConstants
	SigningSecret                       string //Should be moved to some Secret management system like Jana
	MaxUnsuccessfulLoginAttemptsAllowed uint8
	ServerPort                          string
	KafkaBootStrapServer                string
	KafkaLoginTopic                     string
	DBServerURI                         string
}

var instance *Config

//GetInstance method for accessing singleton Config instance
func GetInstance() *Config {

	if instance == nil {
		instance = &Config{
			Constants: &AppConstants{
				HashingPasswordError:    "Error While Hashing Password, Try Again.",
				UserCreationError:       "Error While Creating User, Try Again.",
				RegistrationSuccessful:  "Registration Successful.",
				UsernameExistsError:     "Username already Exists.",
				InvalidCredentialsError: "Invalid credentials.",
				InvalidUsername:         "Invalid username.",
				InvalidPassword:         "Invalid password.",
				AccountLockedError:      "Account is locked.",
				TokenGenerationError:    "Error while generating access token!",
				SuccessfulLogin:         "Successful login",
				MissingAuthTokenError:   "Missing authorization token",
				MissingUsernameError:    "Missing username.",
				MissingPasswordError:    "Missing user password.",
				MissingEmailError:       "Missing user email",
				MissingFirstnameError:   "Missing user firstname",
				InvalidTokenError:       "Invalid access token.",
				SigningError:            "Unexpected signing method",
			},
			SigningSecret:                       "d5bHh9iQwWP4tJtGMJso44QpLyIyUPIKDmTky7DSzJiy8dTLoVltfWCXEqI2Xceb",
			ServerPort:                          "8080",
			KafkaBootStrapServer:                "localhost",
			KafkaLoginTopic:                     "login_service_topic",
			DBServerURI:                         "mongodb://localhost:27017",
			MaxUnsuccessfulLoginAttemptsAllowed: 5,
		}
	}
	return instance
}
