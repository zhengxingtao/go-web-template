package conf

import (
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

func SwaggerConf() *restfulspec.Config {
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	return &config
}

func enrichSwaggerObject(swo *spec.Swagger) {

	//设置请求头中携带 token 属性
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
