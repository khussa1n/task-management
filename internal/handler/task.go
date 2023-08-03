package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"github.com/khussa1n/task-management/internal/entity"
	"log"
	"net/http"
)

func (h *Handler) createTask(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.TaskCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	req.Tasks.UserID = userID
	task, err := h.srvs.CreateTask(ctx, &req.Tasks)
	if err != nil {
		log.Printf("can not create task: %w", err)
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -3,
			Message: "invalid to create task",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data: entity.Tasks{
			ID:           task.ID,
			UserID:       task.UserID,
			CreatedDate:  task.CreatedDate,
			TaskName:     task.TaskName,
			Description:  task.Description,
			StatusID:     task.StatusID,
			DeadlineFrom: task.DeadlineFrom,
			DeadlineTo:   task.DeadlineTo,
			PriorityID:   task.PriorityID,
			ParentTaskID: task.ParentTaskID,
		},
	})
}

func (h *Handler) getAllTasks(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	tasks, err := h.srvs.GetAllTasks(ctx, userID)
	if err != nil {
		log.Printf("can not get all tasks: %w", err)
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -3,
			Message: "invalid to get all tasks",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data:    tasks,
	})
}