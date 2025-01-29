package v1

import (
	"pet/internal/data/ent"
	"pet/internal/dto/request"
	"pet/internal/dto/response"
	"pet/internal/service"
	"pet/pkg/http/gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommunityHandler struct {
	communityService *service.CommunityService
}

func NewCommunityHandler(opt *Option) *CommunityHandler {
	return &CommunityHandler{
		communityService: opt.CommunitySrv,
	}
}

func (h *CommunityHandler) RegisterRoute(r *gin.RouterGroup) {
	communities := r.Group("/communities")
	{
		communities.POST("", h.createCommunity)                 // 创建帖子
		communities.PUT("/:id", h.updateCommunity)              // 更新帖子
		communities.GET("/:id", h.getCommunity)                 // 获取帖子
		communities.DELETE("/:id", h.deleteCommunity)           // 删除帖子
		communities.GET("", h.listCommunity)                    // 帖子列表
		communities.GET("/author/:id", h.listCommunityByAuthor) // 获取用户的帖子
		communities.GET("/pet/:id", h.listCommunityByPet)       // 获取宠物的帖子
		communities.GET("/type/:type", h.listCommunityByType)   // 获取指定类型的帖子
		communities.PUT("/:id/likes", h.updateCommunityLikes)   // 更新点赞数
		communities.PUT("/:id/views", h.updateCommunityViews)   // 更新浏览量
	}
}

// createCommunity 创建帖子
func (h *CommunityHandler) createCommunity(c *gin.Context) {
	var req request.CreateCommunityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	community := &ent.Community{
		AuthorID: req.AuthorID,
		PetID:    req.PetID,
		Title:    req.Title,
		Type:     req.Type,
		Content:  req.Content,
		Images:   req.Images,
		Tags:     req.Tags,
		IsPinned: req.IsPinned,
	}

	result, err := h.communityService.CreateCommunity(c.Request.Context(), community)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    response.NewCommunityResponse(result),
		"message": "success",
	})
}

// updateCommunity 更新帖子
func (h *CommunityHandler) updateCommunity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req request.UpdateCommunityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	community := &ent.Community{
		ID:       id,
		AuthorID: req.AuthorID,
		PetID:    req.PetID,
		Title:    req.Title,
		Type:     req.Type,
		Content:  req.Content,
		Images:   req.Images,
		Tags:     req.Tags,
		IsPinned: req.IsPinned,
	}

	utils.NewResponse(c)(h.communityService.UpdateCommunity(c.Request.Context(), community))
}

// getCommunity 获取帖子
func (h *CommunityHandler) getCommunity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	utils.NewResponse(c)(h.communityService.GetCommunity(c.Request.Context(), id))
}

// deleteCommunity 删除帖子
func (h *CommunityHandler) deleteCommunity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	if err := h.communityService.DeleteCommunity(c.Request.Context(), id); err != nil {
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

// listCommunity 帖子列表
func (h *CommunityHandler) listCommunity(c *gin.Context) {
	var req request.ListCommunityRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	utils.NewResponse(c)(h.communityService.ListCommunity(c.Request.Context(), req.Page, req.PageSize))
}

// listCommunityByAuthor 获取用户的帖子
func (h *CommunityHandler) listCommunityByAuthor(c *gin.Context) {
	authorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid author id",
		})
		return
	}

	utils.NewResponse(c)(h.communityService.ListCommunityByAuthor(c.Request.Context(), authorID))
}

// listCommunityByPet 获取宠物的帖子
func (h *CommunityHandler) listCommunityByPet(c *gin.Context) {
	petID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid pet id",
		})
		return
	}

	utils.NewResponse(c)(h.communityService.ListCommunityByPet(c.Request.Context(), petID))
}

// listCommunityByType 获取指定类型的帖子
func (h *CommunityHandler) listCommunityByType(c *gin.Context) {
	var req request.ListCommunityRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	postType := c.Param("type")
	utils.NewResponse(c)(h.communityService.ListCommunityByType(c.Request.Context(), postType, req.Page, req.PageSize))
}

// updateCommunityLikes 更新点赞数
func (h *CommunityHandler) updateCommunityLikes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req request.UpdateCommunityLikesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if err := h.communityService.UpdateCommunityLikes(c.Request.Context(), id, req.Increment); err != nil {
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

// updateCommunityViews 更新浏览量
func (h *CommunityHandler) updateCommunityViews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	if err := h.communityService.UpdateCommunityViews(c.Request.Context(), id); err != nil {
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
