package response

import (
	"pet/internal/data/ent"
	"time"
)

// CommunityResponse 社区帖子响应
type CommunityResponse struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"author_id"`
	PetID     int       `json:"pet_id,omitempty"`
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	Images    []string  `json:"images,omitempty"`
	Likes     int       `json:"likes"`
	Comments  int       `json:"comments"`
	Views     int       `json:"views"`
	IsPinned  bool      `json:"is_pinned"`
	Tags      []string  `json:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommunityListResponse 社区帖子列表响应
type CommunityListResponse struct {
	Total int64               `json:"total"`
	Items []CommunityResponse `json:"items"`
}

// NewCommunityResponse 将 ent.Community 转换为 CommunityResponse
func NewCommunityResponse(community *ent.Community) *CommunityResponse {
	return &CommunityResponse{
		ID:        community.ID,
		AuthorID:  community.AuthorID,
		PetID:     community.PetID,
		Title:     community.Title,
		Type:      community.Type,
		Content:   community.Content,
		Images:    community.Images,
		Likes:     community.Likes,
		Comments:  community.Comments,
		Views:     community.Views,
		IsPinned:  community.IsPinned,
		Tags:      community.Tags,
		CreatedAt: community.CreatedAt,
		UpdatedAt: community.UpdatedAt,
	}
}

// NewCommunityListResponse 将 []*ent.Community 转换为 CommunityListResponse
func NewCommunityListResponse(communities []*ent.Community, total int64) *CommunityListResponse {
	items := make([]CommunityResponse, len(communities))
	for i, community := range communities {
		items[i] = *NewCommunityResponse(community)
	}
	return &CommunityListResponse{
		Total: total,
		Items: items,
	}
}
