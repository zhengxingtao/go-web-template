package controller

import (
	"github.com/emicklei/go-restful"
	"honor/dto"
	"honor/service"
	"honor/utils"
	"io"
)

func QueryNewsById(req *restful.Request, rep *restful.Response) {
	setResponseHeader(rep)
	id := req.PathParameter("id")
	result := service.QueryNewsById(id)
	success := utils.Success(&result)
	if _, err := io.WriteString(rep, success); err != nil {
		panic(err.Error())
	}
}

func QueryNewsList(req *restful.Request, rep *restful.Response) {
	setResponseHeader(rep)
	criteria := dto.NewsCriteria{}
	if err := req.ReadEntity(&criteria); err != nil {
		paramError := utils.ParamError()
		_, _ = io.WriteString(rep, paramError)
		return
	}
	list := service.QueryNewsList(&criteria)
	success := utils.Success(list)
	if _, err := io.WriteString(rep, success); err != nil {
		panic(err.Error())
	}
}

//统一处理请求头
func setResponseHeader(rep *restful.Response) {
	rep.AddHeader("Content-Type", "application/json;charset=UTF-8")
}
