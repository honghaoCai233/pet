package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:           []string{"*"}, // 允许所有方法
		AllowHeaders:           nil,
		AllowCredentials:       true,
		AllowAllOrigins:        true,
		AllowWildcard:          true,
		MaxAge:                 12 * time.Hour,
		AllowWebSockets:        true,
		AllowBrowserExtensions: true,
		ExposeHeaders:          []string{"*"}, // 暴露所有头
	}
	return cors.New(config)
}
