package response

import (
	"github.com/gogf/gf/net/ghttp"
)

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// GetListResponse 列表带分页数据返回结构
type GetListResponse struct {
	Page      int         `json:"page"`
	Size      int         `json:"size"`
	TotalSize int         `json:"totalSize"`
	Data      interface{} `json:"data"`
}

// Json 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JsonExit 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, msg string, data ...interface{}) {
	Json(r, code, msg, data...)
	r.Exit()
}
