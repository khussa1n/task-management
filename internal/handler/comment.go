package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createComment(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.CommentCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	req.Comments.UserID = userID
	comment, err := h.srvs.CreateComment(ctx, &req.Comments)
	if err != nil {
		log.Printf("can not create task: %w", err)
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -3,
			Message: "invalid to create comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data:    comment,
	})
}

func (h *Handler) getAllComments(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	comments, err := h.srvs.GetAllComments(ctx, userID)
	if err != nil {
		log.Printf("can not get all tasks: %w", err)
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: "invalid to get all comments",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data:    comments,
	})
}

func (h *Handler) deleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("can not get id param: %w", err)
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "invalid id param",
		})
		return
	}

	err = h.srvs.DeleteComment(ctx, int64(id))
	if err != nil {
		log.Printf("can not update task: %w", err)
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -3,
			Message: "invalid to delete comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Code:    0,
		Message: "success",
		Data:    "True",
	})
}
