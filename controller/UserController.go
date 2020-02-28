package controller

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"honor/model"
	"honor/service"
	"honor/utils"
	"io"
)

func CreateUser(req *restful.Request, rep *restful.Response) {
	setResponseHeader(rep)
	user := model.User{}
	if err := req.ReadEntity(&user); err != nil {
		paramError := utils.ParamError()
		_, _ = io.WriteString(rep, paramError)
		return
	}
	fmt.Print(user)
	success := utils.Success(service.CreateUser(&user))
	if _, err := io.WriteString(rep, success); err != nil {
		_, _ = io.WriteString(rep, utils.ServerError())
		return
	}
}

func QueryUserById(req *restful.Request, rep *restful.Response) {
	setResponseHeader(rep)
	uid := req.PathParameter("uid")
	user, err := service.QueryUserById(uid)
	if err != nil {
		_ = rep.WriteErrorString(404, utils.CommonError(404, err.Error()))
		return
	}
	success := utils.Success(user)
	_, _ = io.WriteString(rep, success)
	return
}
