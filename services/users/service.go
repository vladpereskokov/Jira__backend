package users

import (
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "Users"

func Insert(mongo *mgo.Database, user interface{}) (result interface{}, err error) {
	return user, mongo.C(collection).Insert(user)
}

func All(mongo *mgo.Database, _ interface{}) (result interface{}, err error) {
	result = new(models.UsersList)
	err = mongo.C(collection).Find(bson.M{}).All(result)
	return
}

func GetUserByEmailAndPassword(mongo *mgo.Database, user interface{}) (result interface{}, err error) {
	result = new(models.User)
	err = mongo.C(collection).Find(user).One(result)
	return
}
