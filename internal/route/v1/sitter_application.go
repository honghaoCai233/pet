package v1

import (
	"pet/internal/data/ent"
	"pet/internal/dto/request"
	"pet/internal/service"
	"pet/pkg/http/gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SitterApplicationHandler struct {
	sitterApplicationService *service.SitterApplicationService
}

func NewSitterApplicationHandler(opt *Option) *SitterApplicationHandler {
	return &SitterApplicationHandler{
		sitterApplicationService: opt.SitterApplicationSrv,
	}
}

func (h *SitterApplicationHandler) RegisterRoute(r *gin.RouterGroup) {
	applications := r.Group("/sitter-applications")
	{
		applications.POST("", h.createApplication)                 // 创建申请
		applications.PUT("/:id", h.updateApplication)              // 更新申请
		applications.GET("/:id", h.getApplication)                 // 获取申请
		applications.DELETE("/:id", h.deleteApplication)           // 删除申请
		applications.GET("", h.listApplications)                   // 申请列表
		applications.PUT("/:id/status", h.updateApplicationStatus) // 更新申请状态
		applications.GET("/user/:id", h.getApplicationByUser)      // 获取用户的申请
	}
}

// createApplication 创建申请
func (h *SitterApplicationHandler) createApplication(c *gin.Context) {
	var req request.CreateSitterApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entApplication := &ent.SitterApplication{
		Experience:     req.Experience,
		ExpectedSalary: req.ExpectedSalary,
	}

	utils.NewResponse(c)(h.sitterApplicationService.CreateApplication(c.Request.Context(), entApplication))
}

// updateApplication 更新申请
func (h *SitterApplicationHandler) updateApplication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req request.UpdateSitterApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entApplication := &ent.SitterApplication{
		ID:             id,
		Experience:     req.Experience,
		ExpectedSalary: req.ExpectedSalary,
	}

	utils.NewResponse(c)(h.sitterApplicationService.UpdateApplication(c.Request.Context(), entApplication))
}

// getApplication 获取申请
func (h *SitterApplicationHandler) getApplication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	utils.NewResponse(c)(h.sitterApplicationService.GetApplication(c.Request.Context(), id))
}

// deleteApplication 删除申请
func (h *SitterApplicationHandler) deleteApplication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	if err := h.sitterApplicationService.DeleteApplication(c.Request.Context(), id); err != nil {
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
		"message": "success",
	})
}

// listApplications 申请列表
func (h *SitterApplicationHandler) listApplications(c *gin.Context) {
	var req request.ListSitterApplicationsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	utils.NewResponse(c)(h.sitterApplicationService.ListApplications(c.Request.Context(), req.Page, req.PageSize))
}

// updateApplicationStatus 更新申请状态
func (h *SitterApplicationHandler) updateApplicationStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req request.UpdateSitterApplicationStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if err := h.sitterApplicationService.UpdateApplicationStatus(c.Request.Context(), id, req.Status, &req.Message); err != nil {
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
		"message": "success",
	})
}

// getApplicationByUser 获取用户的申请
func (h *SitterApplicationHandler) getApplicationByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid user id",
		})
		return
	}

	utils.NewResponse(c)(h.sitterApplicationService.GetApplication(c.Request.Context(), userID))
}
