package response

import "time"

// AddressResponse 地址响应
type AddressResponse struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Name         string    `json:"name"`
	Province     string    `json:"province"`
	City         string    `json:"city"`
	District     string    `json:"district"`
	Street       string    `json:"street"`
	DetailedInfo string    `json:"detailed_info"`
	ContactName  string    `json:"contact_name"`
	ContactPhone string    `json:"contact_phone"`
	IsDefault    bool      `json:"is_default"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// AddressListResponse 地址列表响应
type AddressListResponse struct {
	Total int               `json:"total"`
	Items []AddressResponse `json:"items"`
}
