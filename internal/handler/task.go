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
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.TaskCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	req.Tasks.UserID = userID
	task, err := h.srvs.CreateTask(ctx, &req.Tasks)
	if err != nil {
		log.Printf("can not create task: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid to create task",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
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
			Message: "can't get user id from auth",
		})
		return
	}

	tasks, err := h.srvs.GetAllTasks(ctx, userID)
	if err != nil {
		log.Printf("can not get all tasks: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid to get all tasks",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Message: "success",
		Data:    tasks,
	})
}

func (h *Handler) updateTask(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.TaskUpdateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil || req.Tasks.ID == 0 {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	req.Tasks.UserID = userID
	task, err := h.srvs.UpdateTask(ctx, &req.Tasks)
	if err != nil {
		log.Printf("can not update task: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid to update task",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Message: "success",
		Data:    task,
	})
}
