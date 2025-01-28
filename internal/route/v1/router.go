package v1

import "github.com/gin-gonic/gin"

type Router interface {
	RegisterRoute(r *gin.RouterGroup)
}

func NewRouters(opt *Option) []Router {
	rv := make([]Router, 0)
	rv = append(rv,
		NewUserHandler(opt),
		NewPetHandler(opt),
		NewTaskHandler(opt),
		NewCommunityHandler(opt),
		NewSitterApplicationHandler(opt),
	)
	return rv
}
