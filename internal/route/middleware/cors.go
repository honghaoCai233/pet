package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{
			"Host",
			"Origin",
			"Content-Length",
			"Content-Type",  // 已覆盖 multipart/form-data
			"Authorization", // 已覆盖 Bearer token
			"X-Response-Time",
			"X-Real-IP",
			// 新增以下请求头
			"Referer",            // 处理 referer 头
			"Sec-CH-UA",          // 客户端提示头
			"Sec-CH-UA-Mobile",   // 客户端提示头
			"Sec-CH-UA-Platform", // 客户端提示头
			"User-Agent",         // 用户代理头
		},
		AllowCredentials:       false,
		MaxAge:                 12 * time.Hour,
		AllowWebSockets:        true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowOrigins:           []string{"http://192.168.1.14:10086"},
		AllowAllOrigins:        false,
	}
	return cors.New(config)
}
