package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"github.com/khussa1n/task-management/internal/entity"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createAction(ctx *gin.Context) {
	var req api.ActionCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	action, err := h.srvs.CreateAction(ctx, &req.Actions)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Message: "success",
		Data: entity.Roles{
			ID:       action.ID,
			RoleName: action.ActionName,
		},
	})
}

func (h *Handler) getAllActions(ctx *gin.Context) {
	action, err := h.srvs.GetAllActions(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.OK{
		Message: "success",
		Data:    action,
	})
}

func (h *Handler) deleteAction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("can not get id: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid id param",
		})
		return
	}

	err = h.srvs.DeleteAction(ctx, int64(id))
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
