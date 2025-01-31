package v1

import (
	"pet/internal/data/ent"
	"pet/internal/service"
	"pet/pkg/http/gin/utils"
	"strconv"

	"pet/internal/dto/request"

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

func (h *AddressHandler) RegisterRoute(r *gin.RouterGroup) {
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
	var req request.CreateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entAddress := &ent.Address{
		UserID:       req.UserID,
		Name:         req.Name,
		Province:     req.Province,
		City:         req.City,
		District:     req.District,
		Street:       req.Street,
		DetailedInfo: req.DetailedInfo,
		ContactName:  req.ContactName,
		ContactPhone: req.ContactPhone,
		IsDefault:    req.IsDefault,
	}

	utils.NewResponse(c)(h.addressService.CreateAddress(c.Request.Context(), entAddress))
}

// updateAddress 更新地址
func (h *AddressHandler) updateAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req request.UpdateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entAddress := &ent.Address{
		ID:           id,
		Name:         req.Name,
		Province:     req.Province,
		City:         req.City,
		District:     req.District,
		Street:       req.Street,
		DetailedInfo: req.DetailedInfo,
		ContactName:  req.ContactName,
		ContactPhone: req.ContactPhone,
		IsDefault:    req.IsDefault,
	}

	utils.NewResponse(c)(h.addressService.UpdateAddress(c.Request.Context(), entAddress))
}

// getAddress 获取地址
func (h *AddressHandler) getAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	utils.NewResponse(c)(h.addressService.GetAddress(c.Request.Context(), id))
}

// deleteAddress 删除地址
func (h *AddressHandler) deleteAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	if err := h.addressService.DeleteAddress(c.Request.Context(), id); err != nil {
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

// listAddresses 地址列表
func (h *AddressHandler) listAddresses(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	utils.NewResponse(c)(h.addressService.ListAddresses(c.Request.Context(), page, pageSize))
}

// listAddressesByUser 获取用户的地址列表
func (h *AddressHandler) listAddressesByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid user_id",
		})
		return
	}

	utils.NewResponse(c)(h.addressService.ListAddressesByUser(c.Request.Context(), userID))
}

// getDefaultAddress 获取用户的默认地址
func (h *AddressHandler) getDefaultAddress(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid user_id",
		})
		return
	}

	utils.NewResponse(c)(h.addressService.GetDefaultAddress(c.Request.Context(), userID))
}

// setDefaultAddress 设置默认地址
func (h *AddressHandler) setDefaultAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid user_id",
		})
		return
	}

	if err := h.addressService.SetDefaultAddress(c.Request.Context(), id, userID); err != nil {
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
