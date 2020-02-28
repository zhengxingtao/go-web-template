package utils

import "gopkg.in/mgo.v2"

var (
	MongoClient *mgo.Session
	err         error
)

func init() {
	MongoClient, err = mgo.Dial(GetYmlProperties().Mongodb.Url)
	if err != nil {
		panic(err.Error())
	}
}
