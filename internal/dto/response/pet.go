package response

import "time"

// PetResponse 宠物响应
type PetResponse struct {
	ID               int       `json:"id"`
	OwnerID          int       `json:"owner_id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	Breed            string    `json:"breed"`
	Age              int       `json:"age"`
	Weight           float64   `json:"weight"`
	Gender           string    `json:"gender"`
	Description      string    `json:"description"`
	CareInstructions string    `json:"care_instructions"`
	Photos           []string  `json:"photos"`
	Vaccinated       bool      `json:"vaccinated"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// PetListResponse 宠物列表响应
type PetListResponse struct {
	Total int           `json:"total"`
	Items []PetResponse `json:"items"`
}
