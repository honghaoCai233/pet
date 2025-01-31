package response

import "time"

// SitterApplicationResponse 保姆申请响应
type SitterApplicationResponse struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	Experience     string    `json:"experience"`
	Certificates   []string  `json:"certificates"`
	AvailableTime  string    `json:"available_time"`
	ServiceArea    string    `json:"service_area"`
	PetTypes       []string  `json:"pet_types"`
	SelfIntro      string    `json:"self_intro"`
	ExpectedSalary float64   `json:"expected_salary"`
	ServiceContent string    `json:"service_content"`
	ContactPhone   string    `json:"contact_phone"`
	EmergencyPhone string    `json:"emergency_phone"`
	IDCardFront    string    `json:"id_card_front"`
	IDCardBack     string    `json:"id_card_back"`
	HealthCert     string    `json:"health_cert"`
	CriminalRecord string    `json:"criminal_record"`
	Status         string    `json:"status"`
	AuditMessage   string    `json:"audit_message,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	// 关联信息
	User *UserResponse `json:"user,omitempty"`
}

// SitterApplicationListResponse 保姆申请列表响应
type SitterApplicationListResponse struct {
	Total int                         `json:"total"`
	Items []SitterApplicationResponse `json:"items"`
}
