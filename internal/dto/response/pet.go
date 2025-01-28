package response

import (
	"pet/internal/data/ent"
	"time"
)

// PetResponse 宠物响应
type PetResponse struct {
	ID               int       `json:"id"`
	OwnerID          int       `json:"owner_id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	Breed            string    `json:"breed,omitempty"`
	Age              int       `json:"age,omitempty"`
	Weight           float64   `json:"weight,omitempty"`
	Gender           string    `json:"gender,omitempty"`
	Description      string    `json:"description,omitempty"`
	CareInstructions string    `json:"care_instructions,omitempty"`
	Photos           []string  `json:"photos,omitempty"`
	Vaccinated       bool      `json:"vaccinated"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// PetListResponse 宠物列表响应
type PetListResponse struct {
	Total int64         `json:"total"`
	Items []PetResponse `json:"items"`
}

// NewPetResponse 将 ent.Pet 转换为 PetResponse
func NewPetResponse(pet *ent.Pet) *PetResponse {
	return &PetResponse{
		ID:               pet.ID,
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
		CreatedAt:        pet.CreatedAt,
		UpdatedAt:        pet.UpdatedAt,
	}
}

// NewPetListResponse 将 []*ent.Pet 转换为 PetListResponse
func NewPetListResponse(pets []*ent.Pet, total int64) *PetListResponse {
	items := make([]PetResponse, len(pets))
	for i, pet := range pets {
		items[i] = *NewPetResponse(pet)
	}
	return &PetListResponse{
		Total: total,
		Items: items,
	}
}
