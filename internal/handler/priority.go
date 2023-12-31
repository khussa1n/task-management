package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"github.com/khussa1n/task-management/internal/entity"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createPriority(ctx *gin.Context) {
	var req api.PriorityCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	priority, err := h.srvs.CreatePriority(ctx, &req.Priorities)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Message: "success",
		Data: entity.Roles{
			ID:       priority.ID,
			RoleName: priority.PriorityName,
		},
	})
}

func (h *Handler) getAllPriorities(ctx *gin.Context) {
	priorities, err := h.srvs.GetAllPriorities(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Message: "success",
		Data:    priorities,
	})
}

func (h *Handler) deletePriority(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("can not get id: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid id param",
		})
		return
	}

	err = h.srvs.DeletePriority(ctx, int64(id))
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
