package v1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"pet/internal/clients"
	"pet/internal/data/ent"
	"pet/internal/service"
	"pet/pkg/http/gin/utils"
	"strconv"
	"strings"
	"time"

	"pet/internal/dto/request"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
	ossClient   *clients.OSSClient
}

func NewUserHandler(opt *Option) *UserHandler {
	return &UserHandler{
		userService: opt.UserSrv,
		ossClient:   opt.OSSClient,
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
		users.POST("/avatar", h.uploadAvatar)        // 上传头像
	}
}

// uploadAvatar 上传头像
func (h *UserHandler) uploadAvatar(c *gin.Context) {
	id := c.Query("userId")
	userId, _ := strconv.Atoi(id)

	// 先获取用户信息，检查是否有旧头像
	user, err := h.userService.GetUser(c.Request.Context(), userId)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": fmt.Sprintf("获取用户信息失败: %v", err),
		})
		return
	}

	// 如果有旧头像，先删除
	if user.Avatar != "" {
		oldObjectKey := user.Avatar
		// 如果存储的是完整URL，需要提取出objectKey
		if strings.HasPrefix(user.Avatar, "http") {
			parsedURL, err := url.Parse(user.Avatar)
			if err != nil {
				log.Printf("解析旧头像URL失败: %v", err)
			} else {
				// 移除开头的斜杠
				oldObjectKey = strings.TrimPrefix(parsedURL.Path, "/")
			}
		}

		if err := h.ossClient.DeleteObject(oldObjectKey); err != nil {
			log.Printf("删除旧头像失败: %v", err)
			// 继续执行，不影响新头像上传
		}
	}

	// 从 multipart/form-data 获取文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "文件上传失败",
		})
		return
	}
	defer file.Close()

	// 读取文件内容
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "读取文件失败",
		})
		return
	}

	// 检查文件类型
	contentType := http.DetectContentType(fileBytes)
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(400, gin.H{
			"success": false,
			"message": "只允许上传图片文件",
		})
		return
	}

	// 生成文件名 - 确保生成的objectKey带有明确的业务前缀
	fileExt := filepath.Ext(header.Filename)
	objectKey := fmt.Sprintf("avatars/%d_%d%s", userId, time.Now().Unix(), fileExt)

	// 上传到OSS
	avatarURL, err := h.ossClient.UploadFileBytes(fileBytes, objectKey, contentType)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": fmt.Sprintf("上传头像失败: %v", err),
		})
		return
	}

	// 更新用户头像 - 存储完整的URL
	userUpdate := &ent.User{
		ID:     userId,
		Avatar: avatarURL,
	}

	if _, err := h.userService.UpdateUser(c.Request.Context(), userUpdate); err != nil {
		// 如果更新用户信息失败，删除已上传的头像
		if delErr := h.ossClient.DeleteObject(objectKey); delErr != nil {
			log.Printf("删除头像失败: %v", delErr)
		}
		c.JSON(500, gin.H{
			"success": false,
			"message": fmt.Sprintf("更新用户头像失败: %v", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"avatarUrl": avatarURL, // 返回给前端可访问的签名URL
		},
		"message": "上传成功",
	})
}

// register 用户注册
func (h *UserHandler) register(c *gin.Context) {
	var req request.RegisterUserRequest
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
	var req request.LoginUserRequest
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

	var req request.UpdateUserPasswordRequest
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

	var req request.UpdateUserStatusRequest
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

	var req request.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
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
