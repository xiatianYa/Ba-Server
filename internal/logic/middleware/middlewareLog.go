package middleware

import (
	"Ba-Server/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/ip2location/ip2location-go"
	"time"
)

func OptionsLogMiddleWare(r *ghttp.Request) {
	// 记录请求方式
	requestMethod := r.Method

	//不记录GET请求
	if requestMethod != "GET" {
		//创建日志模型
		monLogsOperationModel := g.Model("mon_logs_operation")

		// 记录操作时间
		startTime := time.Now()

		// 记录操作 IP
		clientIP := r.GetClientIp()

		//获取IP所属地
		ipAddr := getIpAddr(clientIP)

		// 记录 User-Agent
		userAgent := r.Header.Get("User-Agent")

		// 记录请求 URI
		requestURI := r.RequestURI

		//请求路径
		operationMethod := r.URL.Path

		// 记录请求参数
		requestParams := r.GetBodyString()

		// 执行后续的处理逻辑
		r.Middleware.Next()

		// 计算操作时间
		elapsedTime := time.Since(startTime)

		monLogsOperation := entity.MonLogsOperation{
			Ip:            clientIP,
			IpAddr:        ipAddr,
			UserAgent:     userAgent,
			RequestUri:    requestURI,
			RequestPath:   operationMethod,
			RequestMethod: requestMethod,
			MethodParams:  requestParams,
			UseTime:       int64(elapsedTime),
		}

		_, _ = monLogsOperationModel.Insert(monLogsOperation)
	}

	// 执行后续的处理逻辑
	r.Middleware.Next()
}

func getIpAddr(ip string) string {
	db, _ := ip2location.OpenDB("resource/ip2location/IP2LOCATION-LITE-DB3.BIN")
	defer db.Close()
	results, _ := db.Get_all(ip)
	return results.Country_long + "|" + results.Region + "|" + results.City
}
