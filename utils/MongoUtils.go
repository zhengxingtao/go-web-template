package utils

import "gopkg.in/mgo.v2"

var (
	Mongo *mgo.Session
	err   error
)

func init() {
	Mongo, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err.Error())
	}
}
