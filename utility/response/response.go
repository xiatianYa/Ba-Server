package response

import (
	"Ba-Server/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	var responseData interface{}
	if data != nil {
		responseData = data
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data interface{}) {
	Json(r, code, message, data)
	r.Exit()
}

func dataReturn(r *ghttp.Request, code int, req ...interface{}) *JsonRes {
	var msg string
	var data interface{}
	if len(req) > 0 {
		msg = gconv.String(req[0])
	}
	if len(req) > 1 {
		data = req[1]
	}
	response := &JsonRes{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	r.SetParam("apiReturnRes", response)
	return response
}

// SuccessExit 请求成功
func SuccessExit(r *ghttp.Request, data interface{}) {
	Json(r, consts.SUCCESS, "请求成功", data)
	r.Exit()
}

// ErrorExit Error 请求异常
func ErrorExit(r *ghttp.Request, message string) {
	res := dataReturn(r, consts.ERROR, message)
	r.Response.WriteJsonExit(res)
	r.Exit()
}

// Auth 认证失败
func Auth(r *ghttp.Request) {
	res := dataReturn(r, consts.ERROR, "用户未登录,请先登录")
	r.Response.WriteJsonExit(res)
	r.Exit()
}

// AuthBlack 认证失败 被冻结拉黑
func AuthBlack(r *ghttp.Request) {
	res := dataReturn(r, consts.BACKER, "您的账号被冻结拉黑,请联系管理员")
	r.Response.WriteJsonExit(res)
	r.Exit()
}

// AuthPermission 权限认证失败(没有权限)
func AuthPermission(r *ghttp.Request) {
	res := dataReturn(r, consts.ERROR, "非法权限,你没有权限操作,请联系管理员")
	r.Response.WriteJsonExit(res)
	r.Exit()
}
