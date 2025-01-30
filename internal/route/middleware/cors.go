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
			"Content-Type",
			"Authorization",
			"X-Response-Time",
			"Proxy-Connection",
			"Referer",
			"User-Agent",
		},
		// 关键修正点
		AllowCredentials:       true,
		AllowAllOrigins:        true,
		AllowOrigins:           []string{"*"},
		MaxAge:                 12 * time.Hour,
		AllowWebSockets:        true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type",
		},
	}
	return cors.New(config)
}
