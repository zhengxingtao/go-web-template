package main

import (
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"honor/controller"
	"honor/dto"
	"honor/model"
	"log"
	"net/http"
)

type NewsResource struct {
}

type HonorResource struct {
}

//路由
func (n NewsResource) WebService() *restful.WebService {
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

func (h HonorResource) WebService() *restful.WebService {
	return nil
}

func main() {

	n := NewsResource{}
	restful.DefaultContainer.Add(n.WebService())
	//swagger配置文件
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	//swagger-ui 配置
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/tony/Documents/go_repository/honor/dist"))))
	// 解决跨域问题
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enrichSwaggerObject(swo *spec.Swagger) {

	//设置请求头中携带 授权属性
	swo.SecurityDefinitions = spec.SecurityDefinitions{
		"internalApiKey": spec.APIKeyAuth("token", "header"),
	}
	swo.Security = []map[string][]string{
		{"internalApiKey": {""}},
	}

	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "HonorService",
			Description: "Honor后台服务API",
			Contact: &spec.ContactInfo{
				Name:  "tonyZheng",
				Email: "1198833831@qq.com",
				URL:   "https://www.huojiang.org",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "https://www.huojiang.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "honor",
		Description: "获奖网-API"}}}
}
