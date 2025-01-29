package v1

import (
	"pet/internal/data/ent"
	"pet/internal/service"
	"pet/pkg/http/gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SitterApplicationHandler struct {
	applicationService *service.SitterApplicationService
}

func NewSitterApplicationHandler(opt *Option) *SitterApplicationHandler {
	return &SitterApplicationHandler{
		applicationService: opt.SitterApplicationSrv,
	}
}

func (h *SitterApplicationHandler) RegisterRoute(r *gin.RouterGroup) {
	applications := r.Group("/sitter-applications")
	{
		applications.POST("", h.createApplication)                      // 创建申请
		applications.PUT("/:id", h.updateApplication)                   // 更新申请
		applications.GET("/:id", h.getApplication)                      // 获取申请
		applications.DELETE("/:id", h.deleteApplication)                // 删除申请
		applications.GET("", h.listApplications)                        // 申请列表
		applications.GET("/sitter/:id", h.listApplicationsBySitter)     // 获取照护者的申请列表
		applications.GET("/owner/:id", h.listApplicationsByOwner)       // 获取宠物主人收到的申请列表
		applications.GET("/pet/:id", h.listApplicationsByPet)           // 获取宠物的申请列表
		applications.GET("/status/:status", h.listApplicationsByStatus) // 获取指定状态的申请列表
		applications.PUT("/:id/status", h.updateApplicationStatus)      // 更新申请状态
	}
}

// createApplication 创建申请
func (h *SitterApplicationHandler) createApplication(c *gin.Context) {
	var application struct {
		SitterID       int     `json:"sitter_id" binding:"required"`
		OwnerID        int     `json:"owner_id" binding:"required"`
		PetID          int     `json:"pet_id" binding:"required"`
		Introduction   string  `json:"introduction"`
		Experience     string  `json:"experience"`
		Availability   string  `json:"availability"`
		ExpectedSalary float64 `json:"expected_salary"`
	}

	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entApplication := &ent.SitterApplication{
		SitterID:       application.SitterID,
		OwnerID:        application.OwnerID,
		PetID:          application.PetID,
		Status:         "pending",
		Introduction:   application.Introduction,
		Experience:     application.Experience,
		Availability:   application.Availability,
		ExpectedSalary: application.ExpectedSalary,
	}

	utils.NewResponse(c)(h.applicationService.CreateApplication(c.Request.Context(), entApplication))
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

	var application struct {
		Introduction   string  `json:"introduction"`
		Experience     string  `json:"experience"`
		Availability   string  `json:"availability"`
		ExpectedSalary float64 `json:"expected_salary"`
	}

	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entApplication := &ent.SitterApplication{
		ID:             id,
		Introduction:   application.Introduction,
		Experience:     application.Experience,
		Availability:   application.Availability,
		ExpectedSalary: application.ExpectedSalary,
	}

	utils.NewResponse(c)(h.applicationService.UpdateApplication(c.Request.Context(), entApplication))
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

	utils.NewResponse(c)(h.applicationService.GetApplication(c.Request.Context(), id))
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

	if err := h.applicationService.DeleteApplication(c.Request.Context(), id); err != nil {
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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	utils.NewResponse(c)(h.applicationService.ListApplications(c.Request.Context(), page, pageSize))
}

// listApplicationsBySitter 获取照护者的申请列表
func (h *SitterApplicationHandler) listApplicationsBySitter(c *gin.Context) {
	sitterID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid sitter id",
		})
		return
	}

	utils.NewResponse(c)(h.applicationService.ListApplicationsBySitter(c.Request.Context(), sitterID))
}

// listApplicationsByOwner 获取宠物主人收到的申请列表
func (h *SitterApplicationHandler) listApplicationsByOwner(c *gin.Context) {
	ownerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid owner id",
		})
		return
	}

	utils.NewResponse(c)(h.applicationService.ListApplicationsByOwner(c.Request.Context(), ownerID))
}

// listApplicationsByPet 获取宠物的申请列表
func (h *SitterApplicationHandler) listApplicationsByPet(c *gin.Context) {
	petID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid pet id",
		})
		return
	}

	utils.NewResponse(c)(h.applicationService.ListApplicationsByPet(c.Request.Context(), petID))
}

// listApplicationsByStatus 获取指定状态的申请列表
func (h *SitterApplicationHandler) listApplicationsByStatus(c *gin.Context) {
	status := c.Param("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	utils.NewResponse(c)(h.applicationService.ListApplicationsByStatus(c.Request.Context(), status, page, pageSize))
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

	var req struct {
		Status       string  `json:"status" binding:"required"`
		RejectReason *string `json:"reject_reason"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if err := h.applicationService.UpdateApplicationStatus(c.Request.Context(), id, req.Status, req.RejectReason); err != nil {
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
