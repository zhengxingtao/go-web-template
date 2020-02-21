package model

import "gopkg.in/mgo.v2/bson"

type News struct {
	Id         bson.ObjectId `bson:"_id"`
	Title      string
	NewsType   int
	NewsBody   string
	Images     []*Image
	CreateTime string
	CreateUser string
}
