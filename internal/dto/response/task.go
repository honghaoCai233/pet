package response

import (
	"pet/internal/data/ent"
	"time"
)

// TaskResponse 任务响应
type TaskResponse struct {
	ID               int       `json:"id"`
	PublisherID      int       `json:"publisher_id"`
	PetID            int       `json:"pet_id"`
	SitterID         int       `json:"sitter_id,omitempty"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Reward           float64   `json:"reward"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	Location         string    `json:"location"`
	Requirements     string    `json:"requirements"`
	VisitsCount      int       `json:"visits_count"`
	CareInstructions string    `json:"care_instructions"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	// 关联信息
	Publisher *UserResponse `json:"publisher,omitempty"`
	Pet       *PetResponse  `json:"pet,omitempty"`
	Sitter    *UserResponse `json:"sitter,omitempty"`
}

// TaskListResponse 任务列表响应
type TaskListResponse struct {
	Total int            `json:"total"`
	Items []TaskResponse `json:"items"`
}

// NewTaskResponse 将 ent.Task 转换为 TaskResponse
func NewTaskResponse(task *ent.Task) *TaskResponse {
	return &TaskResponse{
		ID:               task.ID,
		PublisherID:      task.PublisherID,
		PetID:            task.PetID,
		Title:            task.Title,
		Description:      task.Description,
		Reward:           task.Reward,
		StartTime:        task.StartTime,
		EndTime:          task.EndTime,
		Status:           task.Status,
		Location:         task.Location,
		Requirements:     task.Requirements,
		VisitsCount:      task.VisitsCount,
		CareInstructions: task.CareInstructions,
		SitterID:         task.SitterID,
		CreatedAt:        task.CreatedAt,
		UpdatedAt:        task.UpdatedAt,
	}
}

// NewTaskListResponse 将 []*ent.Task 转换为 TaskListResponse
func NewTaskListResponse(tasks []*ent.Task, total int64) *TaskListResponse {
	items := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		items[i] = *NewTaskResponse(task)
	}
	return &TaskListResponse{
		Total: int(total),
		Items: items,
	}
}
