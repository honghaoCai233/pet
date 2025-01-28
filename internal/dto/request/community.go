package request

// CreateCommunityRequest 创建社区帖子请求
type CreateCommunityRequest struct {
	AuthorID int      `json:"author_id" binding:"required"`
	PetID    int      `json:"pet_id"`
	Title    string   `json:"title" binding:"required"`
	Type     string   `json:"type" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Images   []string `json:"images"`
	Tags     []string `json:"tags"`
	IsPinned bool     `json:"is_pinned"`
}

// UpdateCommunityRequest 更新社区帖子请求
type UpdateCommunityRequest struct {
	AuthorID int      `json:"author_id"`
	PetID    int      `json:"pet_id"`
	Title    string   `json:"title"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	Images   []string `json:"images"`
	Tags     []string `json:"tags"`
	IsPinned bool     `json:"is_pinned"`
}

// ListCommunityRequest 社区帖子列表请求
type ListCommunityRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
	Type     string `form:"type"`
}

// UpdateCommunityLikesRequest 更新帖子点赞请求
type UpdateCommunityLikesRequest struct {
	Increment bool `json:"increment" binding:"required"`
}
