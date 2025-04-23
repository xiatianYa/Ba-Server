package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
	"time"
)

func OptionsLogMiddleWare(r *ghttp.Request) {
	// 记录操作时间
	startTime := time.Now()

	// 记录操作 IP
	clientIP := r.GetClientIp()

	// 记录 User-Agent
	userAgent := r.Header.Get("User-Agent")

	// 记录请求 URI
	requestURI := r.RequestURI

	// 记录请求方式
	requestMethod := r.Method

	// 这里假设请求类型和操作方法说明可以根据业务需求自定义
	// 为了示例，简单设置如下
	requestType := "Web Request"
	operationMethod := r.URL.Path
	operationMethodDescription := "Some operation description"

	// 记录请求参数
	requestParams := r.GetParam("")

	// 执行后续的处理逻辑
	r.Middleware.Next()

	// 计算操作时间
	elapsedTime := time.Since(startTime)

	// 记录日志
	log.Printf("操作时间: %s, 操作 IP: %s, User-Agent: %s, 请求 URI: %s, 请求方式: %s, 请求类型: %s, 操作方法: %s, 操作方法说明: %s, 请求参数: %v, 耗时: %s",
		startTime.Format(time.RFC3339), clientIP, userAgent, requestURI, requestMethod, requestType, operationMethod, operationMethodDescription, requestParams, elapsedTime)
}
