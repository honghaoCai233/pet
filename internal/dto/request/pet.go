package request

// CreatePetRequest 创建宠物请求
type CreatePetRequest struct {
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

// UpdatePetRequest 更新宠物请求
type UpdatePetRequest struct {
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

// ListPetsRequest 宠物列表请求
type ListPetsRequest struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=100"`
}
