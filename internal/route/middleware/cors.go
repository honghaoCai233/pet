package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:           []string{"Host", "Origin", "Content-Length", "Content-Type", "Authorization", "X-Response-Time", "X-Real-IP", "Device-Name", "Height", "Os", "Os-Version", "Referer", "Sec-Ch-Ua", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua-Platform", "Sec-Fetch-Site", "User-Agent", "Width"},
		AllowCredentials:       false,
		MaxAge:                 12 * time.Hour,
		AllowWebSockets:        true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowAllOrigins:        true,
	}
	return cors.New(config)
}
