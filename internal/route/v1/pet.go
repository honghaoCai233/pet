package v1

import (
	"pet/internal/data/ent"
	"pet/internal/dto/request"
	"pet/internal/service"
	"pet/pkg/http/gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PetHandler struct {
	petService *service.PetService
}

func NewPetHandler(opt *Option) *PetHandler {
	return &PetHandler{
		petService: opt.PetSrv,
	}
}

func (h *PetHandler) RegisterRoute(r *gin.RouterGroup) {
	pets := r.Group("/pets")
	{
		pets.POST("", h.createPet)                      // 创建宠物
		pets.PUT("/:id", h.updatePet)                   // 更新宠物
		pets.GET("/:id", h.getPet)                      // 获取宠物
		pets.DELETE("/:id", h.deletePet)                // 删除宠物
		pets.GET("", h.listPets)                        // 宠物列表
		pets.GET("/owner/:owner_id", h.listPetsByOwner) // 获取用户的宠物列表
	}
}

// createPet 创建宠物
func (h *PetHandler) createPet(c *gin.Context) {
	var req request.CreatePetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entPet := &ent.Pet{
		OwnerID:          req.OwnerID,
		Name:             req.Name,
		Type:             req.Type,
		Breed:            req.Breed,
		Age:              req.Age,
		Weight:           req.Weight,
		Gender:           req.Gender,
		Description:      req.Description,
		CareInstructions: req.CareInstructions,
		Photos:           req.Photos,
		Vaccinated:       req.Vaccinated,
	}

	utils.NewResponse(c)(h.petService.CreatePet(c.Request.Context(), entPet))
}

// updatePet 更新宠物
func (h *PetHandler) updatePet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req request.UpdatePetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entPet := &ent.Pet{
		ID:               id,
		OwnerID:          req.OwnerID,
		Name:             req.Name,
		Type:             req.Type,
		Breed:            req.Breed,
		Age:              req.Age,
		Weight:           req.Weight,
		Gender:           req.Gender,
		Description:      req.Description,
		CareInstructions: req.CareInstructions,
		Photos:           req.Photos,
		Vaccinated:       req.Vaccinated,
	}

	utils.NewResponse(c)(h.petService.UpdatePet(c.Request.Context(), entPet))
}

// getPet 获取宠物
func (h *PetHandler) getPet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	utils.NewResponse(c)(h.petService.GetPet(c.Request.Context(), id))
}

// deletePet 删除宠物
func (h *PetHandler) deletePet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	if err := h.petService.DeletePet(c.Request.Context(), id); err != nil {
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

// listPets 宠物列表
func (h *PetHandler) listPets(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	utils.NewResponse(c)(h.petService.ListPets(c.Request.Context(), page, pageSize))
}

// listPetsByOwner 获取用户的宠物列表
func (h *PetHandler) listPetsByOwner(c *gin.Context) {
	ownerID, err := strconv.Atoi(c.Param("owner_id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid owner_id",
		})
		return
	}

	utils.NewResponse(c)(h.petService.ListPetsByOwner(c.Request.Context(), ownerID))
}
