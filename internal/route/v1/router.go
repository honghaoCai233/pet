package v1

import "github.com/gin-gonic/gin"

type Router interface {
	RegisteRoute(r *gin.RouterGroup)
}

func NewRouters(opt *WireOptions) []Router {
	rv := make([]Router, 0)
	// rv = append(rv,
	// )
	return rv
}
