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
			Message: "can't get user id from auth",
		})
		return
	}

	var req api.CommentCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: err.Error(),
		})
		return
	}

	req.Comments.UserID = userID
	comment, err := h.srvs.CreateComment(ctx, &req.Comments)
	if err != nil {
		log.Printf("can not create task: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid to create comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Message: "success",
		Data:    comment,
	})
}

func (h *Handler) getAllComments(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can not get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "can't get user id from auth",
		})
		return
	}

	comments, err := h.srvs.GetAllComments(ctx, userID)
	if err != nil {
		log.Printf("can not get all tasks: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid to get all comments",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Message: "success",
		Data:    comments,
	})
}

func (h *Handler) deleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("can not get id param: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid id param",
		})
		return
	}

	err = h.srvs.DeleteComment(ctx, int64(id))
	if err != nil {
		log.Printf("can not update task: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Message: "invalid to delete comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.OK{
		Message: "success",
		Data:    "True",
	})
}
