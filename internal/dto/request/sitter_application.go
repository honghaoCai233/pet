package request

// CreateSitterApplicationRequest 创建保姆申请请求
type CreateSitterApplicationRequest struct {
	UserID         int      `json:"user_id" binding:"required"`
	Experience     string   `json:"experience" binding:"required"`
	Certificates   []string `json:"certificates"`
	AvailableTime  string   `json:"available_time" binding:"required"`
	ServiceArea    string   `json:"service_area" binding:"required"`
	PetTypes       []string `json:"pet_types" binding:"required"`
	SelfIntro      string   `json:"self_intro" binding:"required"`
	ExpectedSalary float64  `json:"expected_salary" binding:"required"`
	ServiceContent string   `json:"service_content" binding:"required"`
	ContactPhone   string   `json:"contact_phone" binding:"required"`
	EmergencyPhone string   `json:"emergency_phone" binding:"required"`
	IDCardFront    string   `json:"id_card_front" binding:"required"`
	IDCardBack     string   `json:"id_card_back" binding:"required"`
	HealthCert     string   `json:"health_cert" binding:"required"`
	CriminalRecord string   `json:"criminal_record" binding:"required"`
}

// UpdateSitterApplicationRequest 更新保姆申请请求
type UpdateSitterApplicationRequest struct {
	Experience     string   `json:"experience"`
	Certificates   []string `json:"certificates"`
	AvailableTime  string   `json:"available_time"`
	ServiceArea    string   `json:"service_area"`
	PetTypes       []string `json:"pet_types"`
	SelfIntro      string   `json:"self_intro"`
	ExpectedSalary float64  `json:"expected_salary"`
	ServiceContent string   `json:"service_content"`
	ContactPhone   string   `json:"contact_phone"`
	EmergencyPhone string   `json:"emergency_phone"`
	IDCardFront    string   `json:"id_card_front"`
	IDCardBack     string   `json:"id_card_back"`
	HealthCert     string   `json:"health_cert"`
	CriminalRecord string   `json:"criminal_record"`
}

// UpdateSitterApplicationStatusRequest 更新保姆申请状态请求
type UpdateSitterApplicationStatusRequest struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message"` // 审核意见
}

// ListSitterApplicationsRequest 保姆申请列表请求
type ListSitterApplicationsRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
	Status   string `form:"status"` // 可选的状态过滤
}
