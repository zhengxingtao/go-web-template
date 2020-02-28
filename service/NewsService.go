package service

import (
	"gopkg.in/mgo.v2"
	bson2 "gopkg.in/mgo.v2/bson"
	"honor/dto"
	"honor/model"
	"honor/utils"
)

var (
	client *mgo.Collection
)

func init() {
	client = utils.MongoClient.DB("honor").C("news")
}

//新增一个对象到mongodb中
func InsertNews(news *model.News) bool {
	err := client.Insert(news)
	if err != nil {
		return false
	}
	return true
}

func QueryNewsList(criteria *dto.NewsCriteria) *dto.NewsListDto {
	iter := client.Find(bson2.M{"newstype": criteria.NewsType}).Sort("_id").
		Skip((criteria.PageNum - 1) * criteria.PageSize).
		Limit(criteria.PageSize).Iter()
	var new model.News
	var news dto.NewsListDto
	for iter.Next(&new) {
		news.NewsList = append(news.NewsList, new)
	}

	if err := iter.Close(); err != nil {
		panic(err.Error())
	}
	return &news
}

func QueryNewsById(id string) model.News {
	hex := bson2.ObjectIdHex(id)
	var new model.News
	if err := client.FindId(hex).One(&new); err != nil {
		panic(err.Error())
	}
	return new
}
