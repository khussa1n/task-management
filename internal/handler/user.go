package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"github.com/khussa1n/task-management/internal/entity"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("can not get id: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "invalid id param",
		})
		return
	}

	user, err := h.srvs.GetUserByID(ctx, int64(id))
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data: entity.Users{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	})
}

func (h *Handler) updateUser(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.UserUpdateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	req.Users.ID = userID
	user, err := h.srvs.UpdateUser(ctx, &req.Users)
	if err != nil {
		log.Printf("can not update user: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -3,
			Message: "invalid to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data: entity.Users{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Password:  user.Password,
		},
	})
	return
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	err := h.srvs.DeleteUser(ctx, userID)
	if err != nil {
		log.Printf("can not delete user: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: "invalid to delete user",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data:    "True",
	})
}
