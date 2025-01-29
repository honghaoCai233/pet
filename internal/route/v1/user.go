package v1

import (
	"pet/internal/data/ent"
	"pet/internal/service"
	"pet/pkg/http/gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(opt *Option) *UserHandler {
	return &UserHandler{
		userService: opt.UserSrv,
	}
}

func (h *UserHandler) RegisterRoute(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		// 用户认证相关
		users.POST("/register", h.register)           // 用户注册
		users.POST("/login", h.login)                 // 用户登录
		users.POST("/:id/password", h.changePassword) // 修改密码
		users.PUT("/:id/status", h.updateStatus)      // 更新用户状态

		// 用户信息相关
		users.PUT("/:id", h.updateUser)    // 更新用户信息
		users.GET("/:id", h.getUser)       // 获取用户信息
		users.DELETE("/:id", h.deleteUser) // 删除用户
		users.GET("", h.listUsers)         // 用户列表

		// 用户查询相关
		users.GET("/phone/:phone", h.getUserByPhone)          // 通过手机号查询
		users.GET("/username/:username", h.getUserByUsername) // 通过用户名查询
		users.GET("/email/:email", h.getUserByEmail)          // 通过邮箱查询
	}
}

// register 用户注册
func (h *UserHandler) register(c *gin.Context) {
	var req struct {
		Username    string  `json:"username" binding:"required"`
		Password    string  `json:"password" binding:"required,min=6"`
		Email       string  `json:"email"`
		Name        string  `json:"name" binding:"required"`
		Address     string  `json:"address"`
		Phone       string  `json:"phone" binding:"required"`
		Age         int     `json:"age" binding:"required"`
		Role        string  `json:"role"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &ent.User{
		Username:    req.Username,
		Password:    req.Password, // 密码加密会在 service 层处理
		Email:       req.Email,
		Name:        req.Name,
		Address:     req.Address,
		Phone:       req.Phone,
		Age:         req.Age,
		Role:        req.Role,
		Description: req.Description,
		Rating:      req.Rating,
	}

	result, err := h.userService.RegisterUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

// login 用户登录
func (h *UserHandler) login(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		Username string `json:"username" binding:"required"` // 可以是用户名或邮箱
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	//user, err := h.userService.LoginUser(c.Request.Context(), req.Username, req.Password)
	//if err != nil {
	//	c.JSON(401, gin.H{
	//		"code":  401,
	//		"data":  nil,
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"data": user,
	//})
	utils.NewResponse(c)(h.userService.LoginUser(ctx, req.Username, req.Password))
}

// changePassword 修改密码
func (h *UserHandler) changePassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.ChangePassword(c.Request.Context(), id, req.OldPassword, req.NewPassword); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "password changed successfully"})
}

// updateStatus 更新用户状态
func (h *UserHandler) updateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required,oneof=active disabled locked"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.UpdateUserStatus(c.Request.Context(), id, req.Status); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "status updated successfully"})
}

// updateUser 更新用户信息
func (h *UserHandler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var req struct {
		Name        string  `json:"name"`
		Address     string  `json:"address"`
		Phone       string  `json:"phone"`
		Age         int     `json:"age"`
		Role        string  `json:"role"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
		Email       string  `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &ent.User{
		ID:          id,
		Name:        req.Name,
		Address:     req.Address,
		Phone:       req.Phone,
		Age:         req.Age,
		Role:        req.Role,
		Description: req.Description,
		Rating:      req.Rating,
		Email:       req.Email,
	}

	result, err := h.userService.UpdateUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

// getUser 获取用户信息
func (h *UserHandler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.userService.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// deleteUser 删除用户
func (h *UserHandler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user deleted successfully"})
}

// listUsers 用户列表
func (h *UserHandler) listUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, err := h.userService.ListUsers(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

// getUserByPhone 通过手机号获取用户
func (h *UserHandler) getUserByPhone(c *gin.Context) {
	phone := c.Param("phone")

	user, err := h.userService.GetUserByPhone(c.Request.Context(), phone)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// getUserByUsername 通过用户名获取用户
func (h *UserHandler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.userService.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// getUserByEmail 通过邮箱获取用户
func (h *UserHandler) getUserByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.userService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
