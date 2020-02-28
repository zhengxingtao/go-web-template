package controller

import (
	"github.com/emicklei/go-restful"
	"honor/dto"
	"honor/service"
	"honor/utils"
	"io"
)

func UsePwAuthorization(req *restful.Request, rep *restful.Response) {
	setResponseHeader(rep)
	pw := dto.LoginParam{}
	if err := req.ReadEntity(&pw); err != nil {
		paramError := utils.CommonError(401, err.Error())
		_ = rep.WriteErrorString(401, paramError)
		return
	}
	result, err := service.Authorize(pw)
	if err != nil {
		paramError := utils.CommonError(403, err.Error())
		_ = rep.WriteErrorString(403, paramError)
		return
	}
	success := utils.Success(result)
	if _, err := io.WriteString(rep, success); err != nil {
		_ = rep.WriteErrorString(500, utils.ServerError())
		return
	}
}
