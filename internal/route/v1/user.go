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
		users.GET("/phone/:phone", h.getUserByPhone) // 通过手机号查询
	}
}

// register 用户注册
func (h *UserHandler) register(c *gin.Context) {
	var req struct {
		Phone       string  `json:"phone" binding:"required"`
		Password    string  `json:"password" binding:"required,min=6"`
		Name        string  `json:"name"`
		Address     string  `json:"address"`
		Age         int     `json:"age"`
		Role        string  `json:"role" default:"pet_sitter"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
		Avatar      string  `json:"avatar"`
		Gender      string  `json:"gender" binding:"omitempty,oneof=male female"`
		Birthday    string  `json:"birthday"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	user := &ent.User{
		Phone:       req.Phone,
		Password:    req.Password,
		Name:        req.Name,
		Address:     req.Address,
		Age:         req.Age,
		Role:        req.Role,
		Description: req.Description,
		Rating:      req.Rating,
		Avatar:      req.Avatar,
		Gender:      req.Gender,
		Birthday:    req.Birthday,
	}

	utils.NewResponse(c)(h.userService.RegisterUser(c.Request.Context(), user))
}

// login 用户登录
func (h *UserHandler) login(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	utils.NewResponse(c)(h.userService.LoginUser(ctx, req.Phone, req.Password))
}

// changePassword 修改密码
func (h *UserHandler) changePassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if err := h.userService.ChangePassword(c.Request.Context(), id, req.OldPassword, req.NewPassword); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"data":    nil,
		"message": "password changed successfully",
	})
}

// updateStatus 更新用户状态
func (h *UserHandler) updateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required,oneof=active disabled locked"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if err := h.userService.UpdateUserStatus(c.Request.Context(), id, req.Status); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"data":    nil,
		"message": "status updated successfully",
	})
}

// updateUser 更新用户信息
func (h *UserHandler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
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
		Avatar      string  `json:"avatar"`
		Gender      string  `json:"gender" binding:"omitempty,oneof=male female"`
		Birthday    string  `json:"birthday"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
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
		Avatar:      req.Avatar,
		Gender:      req.Gender,
		Birthday:    req.Birthday,
	}

	utils.NewResponse(c)(h.userService.UpdateUser(c.Request.Context(), user))
}

// getUser 获取用户信息
func (h *UserHandler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	utils.NewResponse(c)(h.userService.GetUser(c.Request.Context(), id))
}

// deleteUser 删除用户
func (h *UserHandler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    nil,
		"message": "user deleted successfully",
	})
}

// listUsers 用户列表
func (h *UserHandler) listUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	utils.NewResponse(c)(h.userService.ListUsers(c.Request.Context(), page, pageSize))
}

// getUserByPhone 通过手机号获取用户
func (h *UserHandler) getUserByPhone(c *gin.Context) {
	phone := c.Param("phone")

	utils.NewResponse(c)(h.userService.GetUserByPhone(c.Request.Context(), phone))
}
