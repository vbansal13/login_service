package dao

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/vbansal/login_service/config"
	"github.com/vbansal/login_service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//UserDAO structure
type UserDAO struct {
	userDBName string
	db         *mongo.Database
}

var instance *UserDAO

//GetUserDAOInstance method for accessing singleton UserDAO instance
func GetUserDAOInstance() *UserDAO {

	if instance == nil {
		instance = &UserDAO{
			userDBName: "users",
		}
	}
	return instance
}

//Called only once during package inititalization
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetInstance().DBServerURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	userDAO := GetUserDAOInstance()
	userDAO.db = client.Database("LoginService")

}

//FindUser finds a user with given username in the Database, will return error if user not found
func (uDao *UserDAO) FindUser(username string) (*model.UserModel, error) {
	collection := uDao.db.Collection(uDao.userDBName)
	var user model.UserModel
	err := collection.FindOne(context.Background(), bson.D{{"username", username}}).Decode(&user)
	return &user, err
}

//AddNewUser adds a new user to the Database
func (uDao *UserDAO) AddNewUser(user *model.UserModel) error {
	collection := uDao.db.Collection(uDao.userDBName)

	_, err := collection.InsertOne(context.Background(), user)
	return err
}

//SaveUser saves a user to the Databs
func (uDao *UserDAO) SaveUser(user *model.UserModel) error {
	collection := uDao.db.Collection(uDao.userDBName)
	update := bson.M{
		"$set": user,
	}
	doc := collection.FindOneAndUpdate(context.Background(), bson.D{{"username", user.Username}}, update, nil)
	return doc.Err()
}

//UserNotFoundError convenience method for checking if given error is no user found user.
func (uDao *UserDAO) UserNotFoundError(err error) bool {
	return err.Error() == "mongo: no documents in result"
}
