package conf

import "github.com/emicklei/go-restful"

func CorsConf() *restful.CrossOriginResourceSharing {
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	return &cors
}
