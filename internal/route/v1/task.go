package v1

import (
	"pet/internal/data/ent"
	"pet/internal/dto/request"
	"pet/internal/service"
	"pet/pkg/http/gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(opt *Option) *TaskHandler {
	return &TaskHandler{
		taskService: opt.TaskSrv,
	}
}

func (h *TaskHandler) RegisterRoute(r *gin.RouterGroup) {
	tasks := r.Group("/tasks")
	{
		tasks.POST("", h.createTask)                        // 创建任务
		tasks.PUT("/:id", h.updateTask)                     // 更新任务
		tasks.GET("/:id", h.getTask)                        // 获取任务
		tasks.DELETE("/:id", h.deleteTask)                  // 删除任务
		tasks.GET("", h.listTasks)                          // 任务列表
		tasks.GET("/publisher/:id", h.listTasksByPublisher) // 获取用户发布的任务
		tasks.GET("/sitter/:id", h.listTasksBySitter)       // 获取照看者接受的任务
		tasks.GET("/pet/:id", h.listTasksByPet)             // 获取宠物的任务
		tasks.PUT("/:id/status", h.updateTaskStatus)        // 更新任务状态
	}
}

// createTask 创建任务
func (h *TaskHandler) createTask(c *gin.Context) {
	var req request.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entTask := &ent.Task{
		PublisherID:      req.PublisherID,
		PetID:            req.PetID,
		Title:            req.Title,
		Description:      req.Description,
		Reward:           req.Reward,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		Location:         req.Location,
		Requirements:     req.Requirements,
		VisitsCount:      req.VisitsCount,
		CareInstructions: req.CareInstructions,
	}

	utils.NewResponse(c)(h.taskService.CreateTask(c.Request.Context(), entTask))
}

// updateTask 更新任务
func (h *TaskHandler) updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req request.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	entTask := &ent.Task{
		ID:               id,
		PublisherID:      req.PublisherID,
		PetID:            req.PetID,
		Title:            req.Title,
		Description:      req.Description,
		Reward:           req.Reward,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		Location:         req.Location,
		Requirements:     req.Requirements,
		VisitsCount:      req.VisitsCount,
		CareInstructions: req.CareInstructions,
		SitterID:         req.SitterID,
	}

	utils.NewResponse(c)(h.taskService.UpdateTask(c.Request.Context(), entTask))
}

// getTask 获取任务
func (h *TaskHandler) getTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	utils.NewResponse(c)(h.taskService.GetTask(c.Request.Context(), id))
}

// deleteTask 删除任务
func (h *TaskHandler) deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	if err := h.taskService.DeleteTask(c.Request.Context(), id); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    nil,
		"message": "success",
	})
}

// listTasks 任务列表
func (h *TaskHandler) listTasks(c *gin.Context) {
	var req request.ListTasksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	utils.NewResponse(c)(h.taskService.ListTasks(c.Request.Context(), req.Page, req.PageSize))
}

// listTasksByPublisher 获取用户发布的任务列表
func (h *TaskHandler) listTasksByPublisher(c *gin.Context) {
	publisherID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid publisher id",
		})
		return
	}

	utils.NewResponse(c)(h.taskService.ListTasksByPublisher(c.Request.Context(), publisherID))
}

// listTasksBySitter 获取照看者接受的任务列表
func (h *TaskHandler) listTasksBySitter(c *gin.Context) {
	sitterID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid sitter id",
		})
		return
	}

	utils.NewResponse(c)(h.taskService.ListTasksBySitter(c.Request.Context(), sitterID))
}

// listTasksByPet 获取宠物的任务列表
func (h *TaskHandler) listTasksByPet(c *gin.Context) {
	petID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid pet id",
		})
		return
	}

	utils.NewResponse(c)(h.taskService.ListTasksByPet(c.Request.Context(), petID))
}

// updateTaskStatus 更新任务状态
func (h *TaskHandler) updateTaskStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var req request.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if err := h.taskService.UpdateTaskStatus(c.Request.Context(), id, req.Status); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    nil,
		"message": "success",
	})
}
