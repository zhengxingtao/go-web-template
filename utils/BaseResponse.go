package utils

import (
	"encoding/json"
)

//返回值包装对象
type BaseResponse struct {
	StatusCode int
	StatusMsg  string
	Data       interface{}
}

//请求成功
func Success(o interface{}) string {
	response := BaseResponse{
		StatusCode: 200,
		StatusMsg:  "处理成功",
		Data:       o,
	}
	result, err := json.Marshal(response)
	if err != nil {
		return ServerError()
	}
	return string(result)
}

//服务器内部错误
func ServerError() string {
	response := BaseResponse{
		StatusCode: 500,
		StatusMsg:  "服务端错误",
		Data:       nil,
	}
	result, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}
	return string(result)
}

//参数错误
func ParamError() string {
	response := BaseResponse{
		StatusCode: 400,
		StatusMsg:  "参数错误",
		Data:       nil,
	}
	result, err := json.Marshal(response)
	if err != nil {
		return ServerError()
	}
	return string(result)
}
