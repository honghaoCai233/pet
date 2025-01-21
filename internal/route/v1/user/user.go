package user

import (
	"pet/internal/dto/request"
	"pet/pkg/http/gin/utils"

	"github.com/gin-gonic/gin"
)

type User struct {
	opt *Option
}

func NewUser(opt *Option) *User {
	return &User{opt: opt}
}

func (r *User) RegisterRoute(g *gin.RouterGroup) {
	sg := g.Group("/user")
	{
		sg.POST("/create", r.create)
	}
}

func (r *User) create(ctx *gin.Context) {
	var req *request.CreateUserRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		utils.NewResponse(ctx)(ctx, err)
		return
	}
	utils.NewResponse(ctx)(r.opt.UserSrv.Create(ctx.Request.Context(), req))
}
