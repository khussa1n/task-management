package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createEvent(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.EventCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	req.Events.UserID = userID
	event, err := h.srvs.CreateEvent(ctx, &req.Events)
	if err != nil {
		log.Printf("can not create task: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid to create event",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Message: "success",
		Data:    event,
	})
}

func (h *Handler) getAllEventsByTaskID(ctx *gin.Context) {
	taskID, err := strconv.Atoi(ctx.Param("taskID"))
	if err != nil {
		log.Printf("can not get id: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid id param",
		})
		return
	}

	events, err := h.srvs.GetAllEventsByTaskID(ctx, int64(taskID))
	if err != nil {
		log.Printf("can not get all events: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid to get all events",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Message: "success",
		Data:    events,
	})
}
