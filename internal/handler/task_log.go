package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createTaskLog(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.TaskLogCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	req.TaskLogs.UserID = userID
	taskLog, err := h.srvs.CreateTaskLog(ctx, &req.TaskLogs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Code:    0,
		Message: "success",
		Data:    taskLog,
	})
}

func (h *Handler) getAllTaskLogsByTaskID(ctx *gin.Context) {
	taskID, err := strconv.Atoi(ctx.Param("taskID"))
	if err != nil {
		log.Printf("can not get id: %w", err)
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "invalid id param",
		})
		return
	}

	taskLogs, err := h.srvs.GetAllTaskLogsByTaskID(ctx, int64(taskID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Code:    0,
		Message: "success",
		Data:    taskLogs,
	})
}

func (h *Handler) updateTaskLog(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.TaskLogUpdateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	req.TaskLogs.UserID = userID
	taskLog, err := h.srvs.UpdateTaskLog(ctx, &req.TaskLogs)
	if err != nil {
		log.Printf("can not update task_logs: %w", err)
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -3,
			Message: "invalid to update task_logs",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data:    taskLog,
	})
}
