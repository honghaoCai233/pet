package v1

import (
	"pet/internal/data/ent"
	"pet/internal/service"
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

func (h *UserHandler) RegisterRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.POST("", h.createUser)                 // 创建用户
		users.PUT("/:id", h.updateUser)              // 更新用户
		users.GET("/:id", h.getUser)                 // 获取用户
		users.DELETE("/:id", h.deleteUser)           // 删除用户
		users.GET("", h.listUsers)                   // 用户列表
		users.GET("/phone/:phone", h.getUserByPhone) // 通过手机号获取用户
	}
}

// createUser 创建用户
func (h *UserHandler) createUser(c *gin.Context) {
	var user struct {
		Name        string  `json:"name" binding:"required"`
		Address     string  `json:"address" binding:"required"`
		Phone       string  `json:"phone" binding:"required"`
		Age         int     `json:"age" binding:"required"`
		Role        string  `json:"role"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 转换为 ent.User
	entUser := &ent.User{
		Name:        user.Name,
		Address:     user.Address,
		Phone:       user.Phone,
		Age:         user.Age,
		Role:        user.Role,
		Description: user.Description,
		Rating:      user.Rating,
	}

	result, err := h.userService.CreateUser(c.Request.Context(), entUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

// updateUser 更新用户
func (h *UserHandler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var user struct {
		Name        string  `json:"name"`
		Address     string  `json:"address"`
		Phone       string  `json:"phone"`
		Age         int     `json:"age"`
		Role        string  `json:"role"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 转换为 ent.User
	entUser := &ent.User{
		ID:          id,
		Name:        user.Name,
		Address:     user.Address,
		Phone:       user.Phone,
		Age:         user.Age,
		Role:        user.Role,
		Description: user.Description,
		Rating:      user.Rating,
	}

	result, err := h.userService.UpdateUser(c.Request.Context(), entUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

// getUser 获取用户
func (h *UserHandler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.userService.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// listUsers 用户列表
func (h *UserHandler) listUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, err := h.userService.ListUsers(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

// getUserByPhone 通过手机号获取用户
func (h *UserHandler) getUserByPhone(c *gin.Context) {
	phone := c.Param("phone")

	user, err := h.userService.GetUserByPhone(c.Request.Context(), phone)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
