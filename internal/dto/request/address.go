package request

// CreateAddressRequest 创建地址请求
type CreateAddressRequest struct {
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

// UpdateAddressRequest 更新地址请求
type UpdateAddressRequest struct {
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
