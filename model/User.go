package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id         bson.ObjectId `bson:"_id"`
	Password   string
	Phone      string
	Salt       string
	Status     int
	CreateTime string
	CreateUser string
}

func (u *User) SetSalt(salt string) {
	u.Salt = salt
}

func (u *User) SetPassword(pwMd5 string) {
	u.Password = pwMd5
}
