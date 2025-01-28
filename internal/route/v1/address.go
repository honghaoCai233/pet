package v1

import (
	"pet/internal/data/ent"
	"pet/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	addressService *service.AddressService
}

func NewAddressHandler(opt *Option) *AddressHandler {
	return &AddressHandler{
		addressService: opt.AddressSrv,
	}
}

func (h *AddressHandler) RegisterRoutes(r *gin.RouterGroup) {
	addresses := r.Group("/addresses")
	{
		addresses.POST("", h.createAddress)                         // 创建地址
		addresses.PUT("/:id", h.updateAddress)                      // 更新地址
		addresses.GET("/:id", h.getAddress)                         // 获取地址
		addresses.DELETE("/:id", h.deleteAddress)                   // 删除地址
		addresses.GET("", h.listAddresses)                          // 地址列表
		addresses.GET("/user/:user_id", h.listAddressesByUser)      // 获取用户的地址列表
		addresses.GET("/default/:user_id", h.getDefaultAddress)     // 获取用户的默认地址
		addresses.PUT("/:id/default/:user_id", h.setDefaultAddress) // 设置默认地址
	}
}

// createAddress 创建地址
func (h *AddressHandler) createAddress(c *gin.Context) {
	var address struct {
		UserID       int    `json:"user_id"`
		Name         string `json:"name" binding:"required"`
		Province     string `json:"province" binding:"required"`
		City         string `json:"city" binding:"required"`
		District     string `json:"district" binding:"required"`
		Street       string `json:"street" binding:"required"`
		DetailedInfo string `json:"detailed_info" binding:"required"`
		ContactName  string `json:"contact_name" binding:"required"`
		ContactPhone string `json:"contact_phone" binding:"required"`
		IsDefault    bool   `json:"is_default"`
	}

	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	entAddress := &ent.Address{
		UserID:       address.UserID,
		Name:         address.Name,
		Province:     address.Province,
		City:         address.City,
		District:     address.District,
		Street:       address.Street,
		DetailedInfo: address.DetailedInfo,
		ContactName:  address.ContactName,
		ContactPhone: address.ContactPhone,
		IsDefault:    address.IsDefault,
	}

	result, err := h.addressService.CreateAddress(c.Request.Context(), entAddress)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

// updateAddress 更新地址
func (h *AddressHandler) updateAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var address struct {
		Name         string `json:"name"`
		Province     string `json:"province"`
		City         string `json:"city"`
		District     string `json:"district"`
		Street       string `json:"street"`
		DetailedInfo string `json:"detailed_info"`
		ContactName  string `json:"contact_name"`
		ContactPhone string `json:"contact_phone"`
		IsDefault    bool   `json:"is_default"`
	}

	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	entAddress := &ent.Address{
		ID:           id,
		Name:         address.Name,
		Province:     address.Province,
		City:         address.City,
		District:     address.District,
		Street:       address.Street,
		DetailedInfo: address.DetailedInfo,
		ContactName:  address.ContactName,
		ContactPhone: address.ContactPhone,
		IsDefault:    address.IsDefault,
	}

	result, err := h.addressService.UpdateAddress(c.Request.Context(), entAddress)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

// getAddress 获取地址
func (h *AddressHandler) getAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	address, err := h.addressService.GetAddress(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, address)
}

// deleteAddress 删除地址
func (h *AddressHandler) deleteAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := h.addressService.DeleteAddress(c.Request.Context(), id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// listAddresses 地址列表
func (h *AddressHandler) listAddresses(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	addresses, err := h.addressService.ListAddresses(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, addresses)
}

// listAddressesByUser 获取用户的地址列表
func (h *AddressHandler) listAddressesByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user_id"})
		return
	}

	addresses, err := h.addressService.ListAddressesByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, addresses)
}

// getDefaultAddress 获取用户的默认地址
func (h *AddressHandler) getDefaultAddress(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user_id"})
		return
	}

	address, err := h.addressService.GetDefaultAddress(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, address)
}

// setDefaultAddress 设置默认地址
func (h *AddressHandler) setDefaultAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user_id"})
		return
	}

	if err := h.addressService.SetDefaultAddress(c.Request.Context(), id, userID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
