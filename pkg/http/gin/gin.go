package gin

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"

	"pet/pkg/http/gin/middleware"
)

func NewGinEngine(mode string, logger ...io.Writer) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	if len(logger) == 0 {
		logger = append(logger, os.Stdout)
	}
	gin.DefaultWriter = io.MultiWriter(logger...)

	e := gin.New()
	_ = e.SetTrustedProxies(nil)
	e.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.CORS(),
	)
	return e
}
