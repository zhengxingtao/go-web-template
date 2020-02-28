package main

import (
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"honor/conf"
	"honor/controller"
	"honor/dto"
	"honor/model"
	"honor/service"
	"honor/utils"
	"log"
	"net/http"
)

type NewsResource struct {
}
type UserResource struct {
}
type AuthResource struct {
}

//路由
func (n *NewsResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/news")
	ws.Consumes(restful.MIME_JSON, restful.MIME_XML)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"news - 新闻服务"}

	//通过id获取news
	ws.Route(ws.GET("/{id}").To(controller.QueryNewsById).
		//docs
		Doc("通过id获取新闻").
		Param(ws.PathParameter("id", "新闻编号").DataType("string")).
		Writes(model.News{}).
		Returns(200, "OK", model.News{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))
	//获取news列表
	ws.Route(ws.POST("/queryList").To(controller.QueryNewsList).
		Doc("获取新闻列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(dto.NewsCriteria{}))
	//第三个route

	return ws
}

func (user *UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/user")
	ws.Consumes(restful.MIME_JSON, restful.MIME_XML)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"user - 用户服务"}

	ws.Route(ws.POST("").To(controller.CreateUser).
		Doc("新增用户").Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.User{}))

	ws.Route(ws.GET("/{uid}").Filter(service.AuthJWT).To(controller.QueryUserById).
		Doc("获取单个user").
		Param(ws.PathParameter("uid", "用户id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.User{}))
	return ws
}

func (auth *AuthResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/auth")
	ws.Consumes(restful.MIME_JSON, restful.MIME_XML)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"auth - 授权服务"}

	ws.Route(ws.POST("/login").To(controller.UsePwAuthorization).
		Doc("系统登录").
		Metadata(restfulspec.KeyOpenAPITags, tags).Reads(dto.LoginParam{}))
	return ws
}

func main() {

	news := NewsResource{}
	user := UserResource{}
	auth := AuthResource{}
	restful.DefaultContainer.Add(news.WebService())
	restful.DefaultContainer.Add(user.WebService())
	restful.DefaultContainer.Add(auth.WebService())
	//swagger配置文件
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(*conf.SwaggerConf()))
	// 解决跨域问题
	restful.DefaultContainer.Filter(conf.CorsConf().Filter)
	//swagger-ui 配置
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/tony/Documents/go_repository/honor/dist"))))

	log.Fatal(http.ListenAndServe(utils.GetYmlProperties().Server.Port, nil))
}
