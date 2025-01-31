package request

import "time"

// CreateTaskRequest 创建任务请求
type CreateTaskRequest struct {
	PublisherID      int       `json:"publisher_id" binding:"required"`
	PetID            int       `json:"pet_id" binding:"required"`
	Title            string    `json:"title" binding:"required"`
	Description      string    `json:"description" binding:"required"`
	Reward           float64   `json:"reward" binding:"required"`
	StartTime        time.Time `json:"start_time" binding:"required"`
	EndTime          time.Time `json:"end_time" binding:"required"`
	Location         string    `json:"location" binding:"required"`
	Requirements     string    `json:"requirements"`
	VisitsCount      int       `json:"visits_count"`
	CareInstructions string    `json:"care_instructions"`
}

// UpdateTaskRequest 更新任务请求
type UpdateTaskRequest struct {
	PublisherID      int       `json:"publisher_id"`
	PetID            int       `json:"pet_id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Reward           float64   `json:"reward"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	Location         string    `json:"location"`
	Requirements     string    `json:"requirements"`
	VisitsCount      int       `json:"visits_count"`
	CareInstructions string    `json:"care_instructions"`
	SitterID         *int      `json:"sitter_id"`
}

// UpdateTaskStatusRequest 更新任务状态请求
type UpdateTaskStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

// ListTasksRequest 任务列表请求
type ListTasksRequest struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=100"`
}
