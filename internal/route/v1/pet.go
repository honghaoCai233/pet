package v1

import (
	"pet/internal/data/ent"
	"pet/internal/service"
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

func (h *PetHandler) RegisterRoutes(r *gin.RouterGroup) {
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
	var pet struct {
		OwnerID          int      `json:"owner_id" binding:"required"`
		Name             string   `json:"name" binding:"required"`
		Type             string   `json:"type" binding:"required"`
		Breed            string   `json:"breed"`
		Age              int      `json:"age"`
		Weight           float64  `json:"weight"`
		Gender           string   `json:"gender"`
		Description      string   `json:"description"`
		CareInstructions string   `json:"care_instructions"`
		Photos           []string `json:"photos"`
		Vaccinated       bool     `json:"vaccinated"`
	}

	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	entPet := &ent.Pet{
		OwnerID:          pet.OwnerID,
		Name:             pet.Name,
		Type:             pet.Type,
		Breed:            pet.Breed,
		Age:              pet.Age,
		Weight:           pet.Weight,
		Gender:           pet.Gender,
		Description:      pet.Description,
		CareInstructions: pet.CareInstructions,
		Photos:           pet.Photos,
		Vaccinated:       pet.Vaccinated,
	}

	result, err := h.petService.CreatePet(c.Request.Context(), entPet)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

// updatePet 更新宠物
func (h *PetHandler) updatePet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var pet struct {
		OwnerID          int      `json:"owner_id"`
		Name             string   `json:"name"`
		Type             string   `json:"type"`
		Breed            string   `json:"breed"`
		Age              int      `json:"age"`
		Weight           float64  `json:"weight"`
		Gender           string   `json:"gender"`
		Description      string   `json:"description"`
		CareInstructions string   `json:"care_instructions"`
		Photos           []string `json:"photos"`
		Vaccinated       bool     `json:"vaccinated"`
	}

	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	entPet := &ent.Pet{
		ID:               id,
		OwnerID:          pet.OwnerID,
		Name:             pet.Name,
		Type:             pet.Type,
		Breed:            pet.Breed,
		Age:              pet.Age,
		Weight:           pet.Weight,
		Gender:           pet.Gender,
		Description:      pet.Description,
		CareInstructions: pet.CareInstructions,
		Photos:           pet.Photos,
		Vaccinated:       pet.Vaccinated,
	}

	result, err := h.petService.UpdatePet(c.Request.Context(), entPet)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

// getPet 获取宠物
func (h *PetHandler) getPet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	pet, err := h.petService.GetPet(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pet)
}

// deletePet 删除宠物
func (h *PetHandler) deletePet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := h.petService.DeletePet(c.Request.Context(), id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// listPets 宠物列表
func (h *PetHandler) listPets(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	pets, err := h.petService.ListPets(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pets)
}

// listPetsByOwner 获取用户的宠物列表
func (h *PetHandler) listPetsByOwner(c *gin.Context) {
	ownerID, err := strconv.Atoi(c.Param("owner_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid owner_id"})
		return
	}

	pets, err := h.petService.ListPetsByOwner(c.Request.Context(), ownerID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pets)
}
