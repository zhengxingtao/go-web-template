package service

import (
	"errors"
	bson2 "gopkg.in/mgo.v2/bson"
	"honor/dto"
	"honor/model"
	"honor/utils"
)

func init() {
	client = utils.MongoClient.DB("honor").C("user")
}

//创建user
func CreateUser(user *model.User) bool {
	salt := utils.RandNumber()
	user.SetSalt(salt)
	passwordMd5 := utils.PasswordMd5(user.Password, salt)
	user.SetPassword(passwordMd5)
	if err := client.Insert(user); err != nil {
		panic(err.Error())
	}
	return true
}

//修改user
func UpdateUser(user *model.User) {

}

func QueryUserByPhone(login dto.LoginParam) (model.User, error) {
	user := model.User{}
	if err := client.Find(bson2.M{"phone": login.Phone}).One(&user); err != nil {
		return model.User{}, errors.New(login.Phone + " :用户未注册")
	}
	return user, nil
}

//通过id查询user
func QueryUserById(id string) (model.User, error) {
	idHex := bson2.ObjectIdHex(id)
	var user model.User
	if err := client.FindId(idHex).One(&user); err != nil {
		return user, err
	}
	return user, nil
}
