package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"github.com/khussa1n/task-management/internal/entity"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createStatus(ctx *gin.Context) {
	var req api.StatusCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	status, err := h.srvs.CreateStatus(ctx, &req.Statuses)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Message: "success",
		Data: entity.Statuses{
			ID:         status.ID,
			StatusName: status.StatusName,
		},
	})
}

func (h *Handler) getAllStatuses(ctx *gin.Context) {
	statuses, err := h.srvs.GetAllStatuses(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Message: "success",
		Data:    statuses,
	})
}

func (h *Handler) deleteStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("can not get id: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid id param",
		})
		return
	}

	err = h.srvs.DeleteStatus(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Message: "success",
		Data:    "True",
	})
}
